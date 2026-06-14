/**
 * SubElements List Page
 * Display all curriculum subelements with CRUD operations using page-based forms
 */

import React from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  CardActions,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  Add as AddIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useSubelementsByElement } from '@/services/queries/CPQueryService';
import { deleteSubelement } from '@/api/cp';
import { CurriculumSubelement } from '@/shared/types/domain';
import { useMutation } from '@tanstack/react-query';

const SubElementsListPage: React.FC = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const elementId = searchParams.get('elementId') || '';

  const handleAdd = () => {
    navigate(`/dashboard/curriculum/subelements/new?elementId=${elementId}`);
  };

  const handleEdit = (subelementId: string) => {
    navigate(`/dashboard/curriculum/subelements/${subelementId}?elementId=${elementId}`);
  };

  const deleteMutation = useMutation({
    mutationFn: deleteSubelement,
  });

  const handleDelete = async (subelement: CurriculumSubelement) => {
    if (window.confirm(`Apakah Anda yakin ingin menghapus sub-elemen "${subelement.name}"?`)) {
      try {
        await deleteMutation.mutateAsync(subelement.id);
      } catch (error) {
        console.error('Error deleting subelement:', error);
      }
    }
  };

  const { data: subelements = [], isLoading, error } = useSubelementsByElement(elementId);

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return <Alert severity="error">Error loading subelements</Alert>;
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Subelemen</Typography>
        <Button variant="contained" startIcon={<AddIcon />} onClick={handleAdd}>
          Tambah Subelemen
        </Button>
      </Box>

      <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
        {subelements.map((subelement) => (
          <Box sx={{ width: { xs: '100%', sm: '50%', md: '33.33%' } }} key={subelement.id}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  {subelement.name}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Kode: {subelement.code}
                </Typography>
              </CardContent>
              <CardActions>
                <Button
                  size="small"
                  startIcon={<EditIcon />}
                  onClick={() => handleEdit(subelement.id)}
                >
                  Edit
                </Button>
                <Button
                  size="small"
                  color="error"
                  startIcon={<DeleteIcon />}
                  onClick={() => handleDelete(subelement)}
                >
                  Hapus
                </Button>
              </CardActions>
            </Card>
          </Box>
        ))}
      </Box>
    </Box>
  );
};

export default SubElementsListPage;
