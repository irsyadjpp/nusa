/**
 * Academic Year Management Component
 * Advanced component with CRUD operations, validation, and proper state management
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
  Archive as ArchiveIcon,
  CheckCircle as ActivateIcon,
  CalendarToday as CalendarIcon,
  School as SchoolIcon,
  ArrowForward as SelectIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { AcademicYear } from '@/api/academic-foundation';
import { useAcademicYears } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';
import { useMutation } from '@tanstack/react-query';

interface AcademicYearManagementProps {
  schoolId: string;
  onSelectAcademicYear?: (academicYear: AcademicYear) => void;
}

export const AcademicYearManagement = ({ schoolId, onSelectAcademicYear }: AcademicYearManagementProps) => {
  const navigate = useNavigate();
  const { data: academicYears, isLoading, error } = useAcademicYears({ school_id: schoolId });

  // Mutations
  const activateMutation = useMutation({
    mutationFn: academicFoundationApi.activateAcademicYear,
    onSuccess: () => {
      // The page will handle refetching
    },
  });

  const archiveMutation = useMutation({
    mutationFn: academicFoundationApi.archiveAcademicYear,
    onSuccess: () => {
      // The page will handle refetching
    },
  });

  const handleAdd = () => {
    navigate('/dashboard/academic-foundation/academic-years/new');
  };

  const handleEdit = (academicYearId: string) => {
    navigate(`/dashboard/academic-foundation/academic-years/${academicYearId}`);
  };

  const handleActivate = (academicYear: AcademicYear) => {
    activateMutation.mutate(academicYear.id);
  };

  const handleArchive = (academicYear: AcademicYear) => {
    if (window.confirm(`Apakah Anda yakin ingin mengarsipkan tahun ajaran ${academicYear.name}?`)) {
      archiveMutation.mutate(academicYear.id);
    }
  };

  const getStatusColor = (status: string): 'default' | 'success' | 'warning' | 'error' => {
    switch (status) {
      case 'ACTIVE': return 'success';
      case 'ARCHIVED': return 'default';
      case 'DRAFT': return 'warning';
      default: return 'default';
    }
  };

  const getStatusLabel = (status: string): string => {
    switch (status) {
      case 'ACTIVE': return 'Aktif';
      case 'ARCHIVED': return 'Diarsipkan';
      case 'DRAFT': return 'Draft';
      default: return status;
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
        Gagal memuat data tahun ajaran
      </Alert>
    );
  }

  return (
    <Box sx={{ p: 3 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h5" component="h2">
          <SchoolIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
          Manajemen Tahun Ajaran
        </Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={handleAdd}
        >
          Tambah Tahun Ajaran
        </Button>
      </Box>

      {!academicYears || academicYears.length === 0 ? (
        <Box sx={{ textAlign: 'center', py: 8 }}>
          <SchoolIcon sx={{ fontSize: 64, color: 'text.secondary', mb: 2 }} />
          <Typography variant="h6" color="text.secondary" gutterBottom>
            Belum Ada Tahun Ajaran
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Mulai dengan menambahkan tahun ajaran pertama untuk sekolah Anda
          </Typography>
        </Box>
      ) : (
        <Grid container spacing={3}>
          {academicYears.map((academicYear: AcademicYear) => (
            <Grid size={{ xs: 12, md: 6, lg: 4 }} key={academicYear.id}>
              <Card>
                <CardContent>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'start', mb: 2 }}>
                    <Typography variant="h6" component="h3" gutterBottom>
                      {academicYear.name}
                    </Typography>
                    <Chip
                      label={getStatusLabel(academicYear.status)}
                      color={getStatusColor(academicYear.status)}
                      size="small"
                    />
                  </Box>
                  
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
                    <CalendarIcon sx={{ fontSize: 16, mr: 1, color: 'text.secondary' }} />
                    <Typography variant="body2" color="text.secondary">
                      {new Date(academicYear.start_date).toLocaleDateString('id-ID', {
                        day: 'numeric',
                        month: 'long',
                        year: 'numeric',
                      })}
                    </Typography>
                  </Box>
                  
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <CalendarIcon sx={{ fontSize: 16, mr: 1, color: 'text.secondary' }} />
                    <Typography variant="body2" color="text.secondary">
                      {new Date(academicYear.end_date).toLocaleDateString('id-ID', {
                        day: 'numeric',
                        month: 'long',
                        year: 'numeric',
                      })}
                    </Typography>
                  </Box>

                  {academicYear.current_semester_id && (
                    <Typography variant="body2" color="primary" sx={{ mt: 1 }}>
                      Semester aktif tersedia
                    </Typography>
                  )}
                </CardContent>
                
                <Divider />
                
                <CardActions sx={{ justifyContent: 'space-between' }}>
                  {onSelectAcademicYear && (
                    <Button
                      size="small"
                      startIcon={<SelectIcon />}
                      onClick={() => onSelectAcademicYear(academicYear)}
                    >
                      Kelola Semester
                    </Button>
                  )}
                  
                  <Box sx={{ display: 'flex' }}>
                    <Tooltip title="Edit">
                    <IconButton 
                      size="small" 
                      onClick={() => handleEdit(academicYear.id)}
                      disabled={academicYear.status === 'ARCHIVED'}
                    >
                      <EditIcon fontSize="small" />
                    </IconButton>
                  </Tooltip>
                  
                  {academicYear.status === 'DRAFT' && (
                    <Tooltip title="Aktifkan">
                      <IconButton 
                        size="small" 
                        color="success"
                        onClick={() => handleActivate(academicYear)}
                      >
                        <ActivateIcon fontSize="small" />
                      </IconButton>
                    </Tooltip>
                  )}
                  
                  {academicYear.status === 'ACTIVE' && (
                    <Tooltip title="Arsipkan">
                      <IconButton 
                        size="small" 
                        color="default"
                        onClick={() => handleArchive(academicYear)}
                      >
                        <ArchiveIcon fontSize="small" />
                      </IconButton>
                    </Tooltip>
                  )}
                  </Box>
                </CardActions>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}
    </Box>
  );
};