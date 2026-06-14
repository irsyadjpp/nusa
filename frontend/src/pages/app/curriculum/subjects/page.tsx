/**
 * Subjects List Page
 * Display all curriculum subjects with CRUD operations
 */

import React from 'react';
import {
  Box,
  Typography,
  Button,
  Card,
  CardContent,
  CardActions,
  Grid,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  Add as AddIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { useSubjects } from '@/services/queries/CPQueryService';
import { deleteSubject } from '@/api/cp';
import { CurriculumSubject } from '@/shared/types/domain';
import { useMutation, useQueryClient } from '@tanstack/react-query';

const SubjectsListPage: React.FC = () => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const { data: subjects = [], isLoading, error } = useSubjects();

  const handleAdd = () => {
    navigate('/dashboard/curriculum/subjects/new');
  };

  const handleEdit = (subjectId: string) => {
    navigate(`/dashboard/curriculum/subjects/${subjectId}`);
  };

  const deleteMutation = useMutation({
    mutationFn: deleteSubject,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['subjects'] });
    },
  });

  const handleDelete = async (subject: CurriculumSubject) => {
    if (window.confirm(`Apakah Anda yakin ingin menghapus mata pelajaran "${subject.name}"?`)) {
      try {
        await deleteMutation.mutateAsync(subject.id);
      } catch (error) {
        console.error('Error deleting subject:', error);
      }
    }
  };

  const handleViewDetail = (subject: CurriculumSubject) => {
    navigate(`/dashboard/curriculum/subjects/${subject.id}`);
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return <Alert severity="error">Error loading subjects</Alert>;
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4">Mata Pelajaran</Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={handleAdd}
        >
          Tambah Mata Pelajaran
        </Button>
      </Box>

      <Grid container spacing={3}>
        {subjects.map((subject) => (
          <Grid size={{ xs: 12, md: 6, lg: 4 }} key={subject.id}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  {subject.name}
                </Typography>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Kode: {subject.code}
                </Typography>
                {subject.description && (
                  <Typography variant="body2" color="text.secondary">
                    {subject.description}
                  </Typography>
                )}
              </CardContent>
              <CardActions>
                <Button size="small" onClick={() => handleViewDetail(subject)}>
                  Detail
                </Button>
                <Button
                  size="small"
                  startIcon={<EditIcon />}
                  onClick={() => handleEdit(subject.id)}
                >
                  Edit
                </Button>
                <Button
                  size="small"
                  color="error"
                  startIcon={<DeleteIcon />}
                  onClick={() => handleDelete(subject)}
                >
                  Hapus
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default SubjectsListPage;
