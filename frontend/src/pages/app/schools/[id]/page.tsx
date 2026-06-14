import React from 'react';
import {
  Box,
  Button,
  Card,
  CardContent,
  Chip,
  Container,
  Grid,
  Typography,
  Alert,
  CircularProgress,
  Divider,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
  Phone as PhoneIcon,
  Email as EmailIcon,
  LocationOn as LocationIcon,
  Language as LanguageIcon,
} from '@mui/icons-material';
import { useSchool } from '@/services/queries/SchoolsQueryService';
import { useDeleteSchool, useUpdateSchoolStatus } from '@/services/commands/SchoolsCommandService';
import { useNavigate, useParams } from 'react-router-dom';

const SchoolDetail = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [alert, setAlert] = React.useState<{ show: boolean; message: string; severity: 'success' | 'error' }>({
    show: false,
    message: '',
    severity: 'success',
  });

  const { data: school, isLoading, error, refetch } = useSchool(id!);

  const deleteMutation = useDeleteSchool({
    onSuccess: () => {
      setAlert({ show: true, message: 'School deleted successfully', severity: 'success' });
      navigate('/dashboard/schools');
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to delete school', severity: 'error' });
    },
  });

  const statusMutation = useUpdateSchoolStatus({
    onSuccess: () => {
      setAlert({ show: true, message: 'School status updated successfully', severity: 'success' });
      refetch();
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to update school status', severity: 'error' });
    },
  });

  const handleDeleteClick = () => {
    if (school && window.confirm(`Are you sure you want to delete the school "${school.name}"? This action cannot be undone.`)) {
      deleteMutation.mutate(school.id);
    }
  };

  const handleToggleStatus = () => {
    if (school) {
      statusMutation.mutate({ id: school.id, is_active: !school.is_active });
    }
  };

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <CircularProgress />
      </Box>
    );
  }

  if (error || !school) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error" sx={{ mt: 3 }}>
          Failed to load school details
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/schools')}
          sx={{ mb: 2 }}
        >
          Back to Schools
        </Button>
        <Typography variant="h4" component="h1" gutterBottom>
          School Details
        </Typography>
      </Box>

      {alert.show && (
        <Alert
          severity={alert.severity}
          onClose={() => setAlert({ ...alert, show: false })}
          sx={{ mb: 2 }}
        >
          {alert.message}
        </Alert>
      )}

      <Grid container spacing={3}>
        <Grid size={{ xs: 12, md: 8 }}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 3 }}>
                <Box>
                  <Typography variant="h5" gutterBottom>
                    {school.name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Code: {school.code}
                  </Typography>
                </Box>
                <Chip
                  label={school.is_active ? 'Active' : 'Inactive'}
                  color={school.is_active ? 'success' : 'default'}
                />
              </Box>

              <Divider sx={{ my: 2 }} />

              <Grid container spacing={2}>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <LocationIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Address
                      </Typography>
                      <Typography variant="body1">
                        {school.address || '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <LocationIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        City
                      </Typography>
                      <Typography variant="body1">
                        {school.city || '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <LocationIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        State
                      </Typography>
                      <Typography variant="body1">
                        {school.state || '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <LocationIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Country
                      </Typography>
                      <Typography variant="body1">
                        {school.country || '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <LocationIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Postal Code
                      </Typography>
                      <Typography variant="body1">
                        {school.postal_code || '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <PhoneIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Phone
                      </Typography>
                      <Typography variant="body1">
                        {school.phone || '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <EmailIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Email
                      </Typography>
                      <Typography variant="body1">
                        {school.email || '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <LanguageIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Website
                      </Typography>
                      <Typography variant="body1">
                        {school.website ? (
                          <a href={school.website} target="_blank" rel="noopener noreferrer">
                            {school.website}
                          </a>
                        ) : '-'}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
              </Grid>
            </CardContent>
          </Card>
        </Grid>

        <Grid size={{ xs: 12, md: 4 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Actions
              </Typography>
              <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                <Button
                  variant="contained"
                  startIcon={<EditIcon />}
                  onClick={() => navigate(`/dashboard/schools/${school.id}`)}
                >
                  Edit School
                </Button>
                <Button
                  variant="outlined"
                  onClick={handleToggleStatus}
                  disabled={statusMutation.isPending}
                >
                  {school.is_active ? 'Deactivate' : 'Activate'}
                </Button>
                <Button
                  variant="outlined"
                  color="error"
                  startIcon={<DeleteIcon />}
                  onClick={handleDeleteClick}
                  disabled={deleteMutation.isPending}
                >
                  Delete School
                </Button>
              </Box>
            </CardContent>
          </Card>

          <Card sx={{ mt: 2 }}>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Metadata
              </Typography>
              <Box>
                <Typography variant="body2" color="text.secondary">
                  Created At
                </Typography>
                <Typography variant="body1" gutterBottom>
                  {new Date(school.created_at).toLocaleString()}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Updated At
                </Typography>
                <Typography variant="body1">
                  {new Date(school.updated_at).toLocaleString()}
                </Typography>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
};

export default SchoolDetail;
