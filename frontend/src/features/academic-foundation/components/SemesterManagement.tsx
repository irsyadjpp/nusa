/**
 * Semester Management Component
 * Advanced component with CRUD operations and academic year integration
 * Now uses page-based forms instead of modals
 */

import {
  Box,
  Button,
  Typography,
  Card,
  CardContent,
  CardActions,
  Chip,
  IconButton,
  Grid,
  Alert,
  CircularProgress,
  Divider,
  Tooltip,
} from '@mui/material';
import {
  Add as AddIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
  CalendarToday as CalendarIcon,
  Folder as SemesterIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { Semester } from '@/api/academic-foundation';
import { useSemesters } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';
import { useMutation } from '@tanstack/react-query';

interface SemesterManagementProps {
  academicYearId: string;
  academicYearName: string;
}

export const SemesterManagement = ({ academicYearId, academicYearName }: SemesterManagementProps) => {
  const navigate = useNavigate();

  const { data: semesters, isLoading, error, refetch } = useSemesters({ academic_year_id: academicYearId });

  const deleteMutation = useMutation({
    mutationFn: academicFoundationApi.deleteSemester,
    onSuccess: () => {
      refetch();
    },
  });

  const handleAdd = () => {
    navigate(`/dashboard/academic-foundation/semesters/new?academicYearId=${academicYearId}`);
  };

  const handleEdit = (semesterId: string) => {
    navigate(`/dashboard/academic-foundation/semesters/${semesterId}?academicYearId=${academicYearId}`);
  };

  const handleDelete = (semester: Semester) => {
    if (window.confirm(`Apakah Anda yakin ingin menghapus semester ${semester.name}?`)) {
      deleteMutation.mutate(semester.id);
    }
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Alert severity="error" sx={{ m: 2 }}>
        Gagal memuat data semester
      </Alert>
    );
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h5" component="h2">
          <SemesterIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
          Semester - {academicYearName}
        </Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={handleAdd}
          disabled={semesters && semesters.length >= 2}
        >
          {semesters && semesters.length >= 2 ? 'Maksimal 2 Semester' : 'Tambah Semester'}
        </Button>
      </Box>

      {!semesters || semesters.length === 0 ? (
        <Box sx={{ textAlign: 'center', py: 8 }}>
          <SemesterIcon sx={{ fontSize: 64, color: 'text.secondary', mb: 2 }} />
          <Typography variant="h6" color="text.secondary" gutterBottom>
            Belum Ada Semester
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Tambahkan semester untuk tahun ajaran ini
          </Typography>
        </Box>
      ) : (
        <Grid container spacing={3}>
          {semesters
            .sort((a: Semester, b: Semester) => a.sequence_number - b.sequence_number)
            .map((semester: Semester) => (
              <Grid size={{ xs: 12, md: 6 }} key={semester.id}>
                <Card>
                  <CardContent>
                    <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'start', mb: 2 }}>
                      <Box>
                        <Typography variant="h6" component="h3" gutterBottom>
                          {semester.name}
                        </Typography>
                        <Box sx={{ display: 'flex', gap: 1, mt: 1 }}>
                          <Chip
                            label={`Semester ${semester.sequence_number}`}
                            size="small"
                            color="info"
                          />
                          <Chip
                            label={semester.status === 'ACTIVE' ? 'Aktif' : 'Tidak Aktif'}
                            size="small"
                            color={semester.status === 'ACTIVE' ? 'success' : 'default'}
                          />
                        </Box>
                      </Box>
                    </Box>
                    
                    <Divider sx={{ my: 2 }} />
                    
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                      <CalendarIcon sx={{ fontSize: 16, mr: 1, color: 'text.secondary' }} />
                      <Typography variant="body2" color="text.secondary">
                        {new Date(semester.start_date).toLocaleDateString('id-ID', {
                          day: 'numeric',
                          month: 'long',
                          year: 'numeric',
                        })}
                      </Typography>
                    </Box>
                    
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <CalendarIcon sx={{ fontSize: 16, mr: 1, color: 'text.secondary' }} />
                      <Typography variant="body2" color="text.secondary">
                        {new Date(semester.end_date).toLocaleDateString('id-ID', {
                          day: 'numeric',
                          month: 'long',
                          year: 'numeric',
                        })}
                      </Typography>
                    </Box>
                  </CardContent>
                  
                  <Divider />
                  
                  <CardActions sx={{ justifyContent: 'flex-end' }}>
                    <Tooltip title="Edit">
                      <IconButton 
                        size="small" 
                        onClick={() => handleEdit(semester.id)}
                      >
                        <EditIcon fontSize="small" />
                      </IconButton>
                    </Tooltip>
                    
                    <Tooltip title="Hapus">
                      <IconButton 
                        size="small" 
                        color="error"
                        onClick={() => handleDelete(semester)}
                      >
                        <DeleteIcon fontSize="small" />
                      </IconButton>
                    </Tooltip>
                  </CardActions>
                </Card>
              </Grid>
            ))}
        </Grid>
      )}
    </Box>
  );
};