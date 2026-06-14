/**
 * Academic Years List Page
 * Display all academic years with CRUD operations using page-based forms
 */

import {
  Box,
  Button,
  Typography,
  CircularProgress,
  Tooltip,
  IconButton,
  Card,
  CardContent,
  CardActions,
  Chip,
  Divider,
} from '@mui/material';
import {
  Add as AddIcon,
  Edit as EditIcon,
  Archive as ArchiveIcon,
  CheckCircle as ActivateIcon,
  CalendarToday as CalendarIcon,
  School as SchoolIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { AcademicYear } from '@/api/academic-foundation';
import { useAcademicYears } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';

const AcademicYearsListPage = () => {
  const navigate = useNavigate();

  const { data: academicYears, isLoading, refetch } = useAcademicYears({
    school_id: 'default-school-id'
  });

  const handleActivate = async (academicYear: AcademicYear) => {
    try {
      await academicFoundationApi.activateAcademicYear(academicYear.id);
      refetch();
    } catch (error) {
      console.error('Error activating academic year:', error);
    }
  };

  const handleArchive = async (academicYear: AcademicYear) => {
    try {
      await academicFoundationApi.archiveAcademicYear(academicYear.id);
      refetch();
    } catch (error) {
      console.error('Error archiving academic year:', error);
    }
  };

  const handleAdd = () => {
    navigate('/dashboard/academic-foundation/academic-years/new');
  };

  const handleEdit = (academicYearId: string) => {
    navigate(`/dashboard/academic-foundation/academic-years/${academicYearId}`);
  };

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
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
          <CalendarIcon sx={{ fontSize: 64, color: 'text.secondary', mb: 2 }} />
          <Typography variant="h6" color="text.secondary" gutterBottom>
            Belum Ada Tahun Ajaran
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Mulai dengan menambahkan tahun ajaran untuk sekolah
          </Typography>
        </Box>
      ) : (
        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
          {academicYears.map((academicYear: any) => (
            <Box sx={{ width: { xs: '100%', md: '50%', lg: '33.33%' } }} key={academicYear.id}>
              <Card>
                <CardContent>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                    <Typography variant="h6" component="h3">
                      {academicYear.name}
                    </Typography>
                    {academicYear.is_current && (
                      <Chip
                        label="Aktif"
                        size="small"
                        color="success"
                      />
                    )}
                  </Box>
                  
                  <Box sx={{ mb: 2 }}>
                    <Typography variant="body2" color="text.secondary">
                      <CalendarIcon sx={{ fontSize: 14, mr: 1, verticalAlign: 'middle' }} />
                      {new Date(academicYear.start_date).toLocaleDateString('id-ID', { 
                        year: 'numeric', 
                        month: 'long', 
                        day: 'numeric' 
                      })} - {new Date(academicYear.end_date).toLocaleDateString('id-ID', { 
                        year: 'numeric', 
                        month: 'long', 
                        day: 'numeric' 
                      })}
                    </Typography>
                  </Box>

                  {!academicYear.is_active && (
                    <Chip label="Nonaktif" size="small" color="default" />
                  )}
                </CardContent>
                <Divider />
                <CardActions>
                  <Tooltip title="Edit">
                    <IconButton size="small" onClick={() => handleEdit(academicYear.id)}>
                      <EditIcon fontSize="small" />
                    </IconButton>
                  </Tooltip>
                  
                  {!academicYear.is_current && academicYear.is_active && (
                    <Tooltip title="Set sebagai aktif">
                      <IconButton size="small" onClick={() => handleActivate(academicYear)}>
                        <ActivateIcon fontSize="small" />
                      </IconButton>
                    </Tooltip>
                  )}
                  
                  {academicYear.is_active && (
                    <Tooltip title="Arsipkan">
                      <IconButton size="small" onClick={() => handleArchive(academicYear)}>
                        <ArchiveIcon fontSize="small" />
                      </IconButton>
                    </Tooltip>
                  )}
                </CardActions>
              </Card>
            </Box>
          ))}
        </Box>
      )}
    </Box>
  );
};

export default AcademicYearsListPage;
