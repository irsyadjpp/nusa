/**
 * Subject Category Form Page
 * Separate page for creating and editing subject categories
 */

import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  TextField,
  Alert,
  CircularProgress,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
  Container,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useNavigate, useParams } from 'react-router-dom';
import { CreateSubjectCategoryRequest } from '@/api/academic-foundation';
import { useSubjectCategories } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';

const SubjectCategoryFormPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id?: string }>();
  const query = new URLSearchParams(window.location.search);
  const parentId = query.get('parentId') || '';

  const [formData, setFormData] = useState<CreateSubjectCategoryRequest>({
    code: '',
    name: '',
    name_indonesian: '',
    description: '',
    parent_id: parentId,
    level: 1,
    sort_order: 1,
    kurikulum_version: '2024',
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);
  const [parentCategory, setParentCategory] = useState<string>(parentId);

  const { data: categories, isLoading } = useSubjectCategories({
    kurikulum_version: '2024'
  });

  // Load category data for editing
  useEffect(() => {
    if (id && categories) {
      const category = categories.find((c: any) => c.id === id);
      if (category) {
        setFormData({
          code: category.code,
          name: category.name,
          name_indonesian: category.name_indonesian,
          description: category.description || '',
          parent_id: category.parent_id || '',
          level: category.level,
          sort_order: category.sort_order,
          kurikulum_version: category.kurikulum_version,
        });
        setParentCategory(category.parent_id || '');
      }
    }
  }, [id, categories]);

  // Calculate level based on parent
  useEffect(() => {
    if (!id && parentCategory && categories) {
      const parentLevel = categories.find((c: any) => c.id === parentCategory)?.level || 0;
      setFormData(prev => ({
        ...prev,
        level: parentLevel + 1,
        parent_id: parentCategory,
      }));
    }
  }, [parentCategory, categories, id]);

  const validateForm = (): boolean => {
    if (!formData.code.trim()) {
      setFormError('Kode kategori harus diisi');
      return false;
    }
    if (!formData.name.trim()) {
      setFormError('Nama kategori (Inggris) harus diisi');
      return false;
    }
    if (!formData.name_indonesian.trim()) {
      setFormError('Nama kategori (Indonesia) harus diisi');
      return false;
    }
    return true;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateForm()) {
      return;
    }

    setSubmitting(true);
    setFormError('');

    try {
      if (id) {
        // Update existing category
        await academicFoundationApi.updateSubjectCategory(id, {
          code: formData.code,
          name: formData.name,
          name_indonesian: formData.name_indonesian,
          description: formData.description,
          parent_id: parentCategory,
          level: formData.level,
          sort_order: formData.sort_order,
        });
      } else {
        // Create new category
        await academicFoundationApi.createSubjectCategory({
          ...formData,
          parent_id: parentCategory,
        });
      }
      navigate('/dashboard/academic-foundation/subject-categories');
    } catch (error: any) {
      setFormError(error.response?.data?.message || 'Gagal menyimpan kategori');
    } finally {
      setSubmitting(false);
    }
  };

  if (isLoading) {
    return (
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/academic-foundation/subject-categories')}
          sx={{ mb: 2 }}
        >
          Kembali ke Daftar Kategori
        </Button>
        <Typography variant="h4" component="h1">
          {id ? 'Edit Kategori Mata Pelajaran' : 'Tambah Kategori Mata Pelajaran Baru'}
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
              {formError && (
                <Alert severity="error" onClose={() => setFormError('')}>
                  {formError}
                </Alert>
              )}
              
              <TextField
                label="Kode Kategori"
                fullWidth
                value={formData.code}
                onChange={(e) => setFormData({ ...formData, code: e.target.value })}
                placeholder="Contoh: MTK, IPA, IPS"
                error={!formData.code.trim()}
                helperText={!formData.code.trim() ? 'Kode kategori harus diisi' : ''}
                required
              />
              
              <TextField
                label="Nama Kategori (Indonesia)"
                fullWidth
                value={formData.name_indonesian}
                onChange={(e) => setFormData({ ...formData, name_indonesian: e.target.value })}
                placeholder="Contoh: Matematika"
                error={!formData.name_indonesian.trim()}
                helperText={!formData.name_indonesian.trim() ? 'Nama kategori harus diisi' : ''}
                required
              />
              
              <TextField
                label="Nama Kategori (Inggris)"
                fullWidth
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Contoh: Mathematics"
                error={!formData.name.trim()}
                helperText={!formData.name.trim() ? 'Nama kategori harus diisi' : ''}
                required
              />
              
              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={3}
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                placeholder="Deskripsi singkat tentang kategori ini"
              />
              
              <FormControl fullWidth>
                <InputLabel>Parent Kategori</InputLabel>
                <Select
                  value={parentCategory}
                  label="Parent Kategori"
                  onChange={(e: SelectChangeEvent) => {
                    setParentCategory(e.target.value);
                    const parentLevel = categories && categories.find((c: any) => c.id === e.target.value)?.level || 0;
                    setFormData({
                      ...formData,
                      parent_id: e.target.value,
                      level: parentLevel + 1,
                    });
                  }}
                >
                  <MenuItem value="">Tidak ada (Root)</MenuItem>
                  {categories?.map((category: any) => (
                    <MenuItem key={category.id} value={category.id}>
                      {category.code} - {category.name_indonesian} (Level {category.level})
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
              
              <TextField
                label="Level"
                fullWidth
                type="number"
                value={formData.level}
                onChange={(e) => setFormData({ ...formData, level: parseInt(e.target.value) })}
                disabled
                helperText="Level akan dihitung otomatis berdasarkan parent kategori"
              />
              
              <TextField
                label="Urutan Sort"
                fullWidth
                type="number"
                value={formData.sort_order}
                onChange={(e) => setFormData({ ...formData, sort_order: parseInt(e.target.value) })}
              />

              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end', mt: 2 }}>
                <Button
                  variant="outlined"
                  onClick={() => navigate('/dashboard/academic-foundation/subject-categories')}
                >
                  Batal
                </Button>
                <Button
                  type="submit"
                  variant="contained"
                  startIcon={<SaveIcon />}
                  disabled={submitting}
                >
                  {submitting ? 'Menyimpan...' : 'Simpan'}
                </Button>
              </Box>
            </Box>
          </form>
        </CardContent>
      </Card>
    </Container>
  );
};

export default SubjectCategoryFormPage;
