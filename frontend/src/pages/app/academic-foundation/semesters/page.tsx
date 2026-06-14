/**
 * Semesters List Page
 * Display all semesters for an academic year with CRUD operations using page-based forms
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
  Delete as DeleteIcon,
  CalendarToday as CalendarIcon,
  Folder as SemesterIcon,
} from '@mui/icons-material';
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { useSemesters } from '@/services/queries/AcademicFoundationQueryService';
import * as academicFoundationApi from '@/api/academic-foundation';

const SemestersListPage = () => {
  const navigate = useNavigate();
  const { academicYearId } = useParams<{ academicYearId: string }>();
  const [searchParams] = useSearchParams();
  const academicYearName = searchParams.get('academicYearName') || '';

  const { data: semesters, isLoading, refetch } = useSemesters({
    academic_year_id: academicYearId!
  });

  const handleDelete = async (semester: any) => {
    if (window.confirm(`Apakah Anda yakin ingin menghapus semester "${semester.name}"?`)) {
      try {
        await academicFoundationApi.deleteSemester(semester.id);
        refetch();
      } catch (error) {
        console.error('Error deleting semester:', error);
      }
    }
  };

  const handleAdd = () => {
    navigate(`/dashboard/academic-foundation/semesters/new?academicYearId=${academicYearId}`);
  };

  const handleEdit = (semesterId: string) => {
    navigate(`/dashboard/academic-foundation/semesters/${semesterId}?academicYearId=${academicYearId}`);
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
          <SemesterIcon sx={{ mr: 1, verticalAlign: 'middle' }} />
          Manajemen Semester - {academicYearName}
        </Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={handleAdd}
        >
          Tambah Semester
        </Button>
      </Box>

      {!semesters || semesters.length === 0 ? (
        <Box sx={{ textAlign: 'center', py: 8 }}>
          <SemesterIcon sx={{ fontSize: 64, color: 'text.secondary', mb: 2 }} />
          <Typography variant="h6" color="text.secondary" gutterBottom>
            Belum Ada Semester
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Mulai dengan menambahkan semester untuk tahun ajaran ini
          </Typography>
        </Box>
      ) : (
        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
          {semesters.map((semester: any) => (
            <Box sx={{ width: { xs: '100%', md: '50%', lg: '33.33%' } }} key={semester.id}>
              <Card>
                <CardContent>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                    <Typography variant="h6" component="h3">
                      {semester.name}
                    </Typography>
                    <Chip
                      label={`Urutan ${semester.sequence}`}
                      size="small"
                      color="info"
                    />
                  </Box>
                  
                  <Box sx={{ mb: 2 }}>
                    <Typography variant="body2" color="text.secondary">
                      <CalendarIcon sx={{ fontSize: 14, mr: 1, verticalAlign: 'middle' }} />
                      {new Date(semester.start_date).toLocaleDateString('id-ID', { 
                        year: 'numeric', 
                        month: 'long', 
                        day: 'numeric' 
                      })} - {new Date(semester.end_date).toLocaleDateString('id-ID', { 
                        year: 'numeric', 
                        month: 'long', 
                        day: 'numeric' 
                      })}
                    </Typography>
                  </Box>

                  {!semester.is_active && (
                    <Chip label="Nonaktif" size="small" color="default" />
                  )}
                </CardContent>
                <Divider />
                <CardActions>
                  <Tooltip title="Edit">
                    <IconButton size="small" onClick={() => handleEdit(semester.id)}>
                      <EditIcon fontSize="small" />
                    </IconButton>
                  </Tooltip>
                  
                  <Tooltip title="Hapus">
                    <IconButton size="small" color="error" onClick={() => handleDelete(semester)}>
                      <DeleteIcon fontSize="small" />
                    </IconButton>
                  </Tooltip>
                </CardActions>
              </Card>
            </Box>
          ))}
        </Box>
      )}
    </Box>
  );
};

export default SemestersListPage;
