/**
 * CP Create Page
 * Create a new Curriculum Plan
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Save as SaveIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { createCP, getSubjects, getPhases } from '@/api/cp';
import { useMutation, useQuery } from '@tanstack/react-query';

const CPCreatePage: React.FC = () => {
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    subject_id: '',
    phase_id: '',
    element_id: '',
    subelement_id: '',
    code: '',
    description: '',
    competency_code: '',
    learning_objectives: '{}',
    competency_standards: '{}',
    time_allocation_hours: '',
    hours_per_week: '',
    version: '1.0',
  });

  const { data: subjects = [] } = useQuery({
    queryKey: ['subjects'],
    queryFn: getSubjects,
  });

  const { data: phases = [] } = useQuery({
    queryKey: ['phases'],
    queryFn: getPhases,
  });

  const createMutation = useMutation({
    mutationFn: createCP,
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const data = {
        subject_id: formData.subject_id,
        phase_id: formData.phase_id,
        element_id: formData.element_id,
        subelement_id: formData.subelement_id,
        code: formData.code,
        description: formData.description,
        competency_code: formData.competency_code || undefined,
        learning_objectives: JSON.parse(formData.learning_objectives),
        competency_standards: JSON.parse(formData.competency_standards),
        time_allocation_hours: parseInt(formData.time_allocation_hours),
        hours_per_week: parseInt(formData.hours_per_week),
        version: formData.version,
      };

      const result = await createMutation.mutateAsync(data);
      navigate(`/cp/${result.id}`);
    } catch (error) {
      alert('Invalid JSON in learning objectives or competency standards');
    }
  };

  return (
    <Box sx={{ p: 3, maxWidth: 800, margin: '0 auto' }}>
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/cp')}
          sx={{ mr: 2 }}
        >
          Kembali
        </Button>
        <Typography variant="h4" component="h1">
          Buat CP Baru
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
              <FormControl fullWidth>
                <InputLabel>Subject</InputLabel>
                <Select
                  value={formData.subject_id}
                  label="Subject"
                  onChange={(e) => setFormData({ ...formData, subject_id: e.target.value })}
                  required
                >
                  {subjects.map((subject) => (
                    <MenuItem key={subject.id} value={subject.id}>
                      {subject.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>

              <FormControl fullWidth>
                <InputLabel>Phase</InputLabel>
                <Select
                  value={formData.phase_id}
                  label="Phase"
                  onChange={(e) => setFormData({ ...formData, phase_id: e.target.value })}
                  required
                >
                  {phases.map((phase) => (
                    <MenuItem key={phase.id} value={phase.id}>
                      {phase.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>

              <TextField
                label="Element ID"
                fullWidth
                variant="outlined"
                value={formData.element_id}
                onChange={(e) => setFormData({ ...formData, element_id: e.target.value })}
                required
              />

              <TextField
                label="Subelement ID"
                fullWidth
                variant="outlined"
                value={formData.subelement_id}
                onChange={(e) => setFormData({ ...formData, subelement_id: e.target.value })}
                required
              />

              <TextField
                label="Kode"
                fullWidth
                variant="outlined"
                value={formData.code}
                onChange={(e) => setFormData({ ...formData, code: e.target.value })}
                required
              />

              <TextField
                label="Deskripsi"
                fullWidth
                multiline
                rows={4}
                variant="outlined"
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                required
              />

              <TextField
                label="Kompetensi Kode"
                fullWidth
                variant="outlined"
                value={formData.competency_code}
                onChange={(e) => setFormData({ ...formData, competency_code: e.target.value })}
              />

              <TextField
                label="Learning Objectives (JSON)"
                fullWidth
                multiline
                rows={4}
                variant="outlined"
                value={formData.learning_objectives}
                onChange={(e) => setFormData({ ...formData, learning_objectives: e.target.value })}
                required
              />

              <TextField
                label="Competency Standards (JSON)"
                fullWidth
                multiline
                rows={4}
                variant="outlined"
                value={formData.competency_standards}
                onChange={(e) => setFormData({ ...formData, competency_standards: e.target.value })}
                required
              />

              <TextField
                label="Alokasi Waktu (Jam)"
                fullWidth
                type="number"
                variant="outlined"
                value={formData.time_allocation_hours}
                onChange={(e) => setFormData({ ...formData, time_allocation_hours: e.target.value })}
                required
              />

              <TextField
                label="Jam per Minggu"
                fullWidth
                type="number"
                variant="outlined"
                value={formData.hours_per_week}
                onChange={(e) => setFormData({ ...formData, hours_per_week: e.target.value })}
                required
              />

              <TextField
                label="Versi"
                fullWidth
                variant="outlined"
                value={formData.version}
                onChange={(e) => setFormData({ ...formData, version: e.target.value })}
                required
              />

              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                <Button variant="outlined" onClick={() => navigate('/cp')}>
                  Batal
                </Button>
                <Button
                  type="submit"
                  variant="contained"
                  startIcon={<SaveIcon />}
                  disabled={createMutation.isPending}
                >
                  {createMutation.isPending ? 'Menyimpan...' : 'Simpan'}
                </Button>
              </Box>
            </Box>
          </form>
        </CardContent>
      </Card>
    </Box>
  );
};

export default CPCreatePage;
