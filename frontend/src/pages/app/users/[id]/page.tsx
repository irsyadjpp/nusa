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
  Avatar,
} from '@mui/material';
import {
  ArrowBack as ArrowBackIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
  Email as EmailIcon,
  Person as PersonIcon,
  School as SchoolIcon,
  Badge as BadgeIcon,
} from '@mui/icons-material';
import { useUser } from '@/services/queries/UsersQueryService';
import { useDeleteUser, useUpdateUserStatus } from '@/services/commands/UsersCommandService';
import { useNavigate, useParams } from 'react-router-dom';

const UserDetail = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const [alert, setAlert] = React.useState<{ show: boolean; message: string; severity: 'success' | 'error' }>({
    show: false,
    message: '',
    severity: 'success',
  });

  const { data: user, isLoading, error, refetch } = useUser(id!);

  const deleteMutation = useDeleteUser({
    onSuccess: () => {
      setAlert({ show: true, message: 'User deleted successfully', severity: 'success' });
      navigate('/dashboard/users');
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to delete user', severity: 'error' });
    },
  });

  const statusMutation = useUpdateUserStatus({
    onSuccess: () => {
      setAlert({ show: true, message: 'User status updated successfully', severity: 'success' });
      refetch();
    },
    onError: (error: any) => {
      setAlert({ show: true, message: error.message || 'Failed to update user status', severity: 'error' });
    },
  });

  const handleDeleteClick = () => {
    if (user && window.confirm(`Are you sure you want to delete the user "${user.name}"? This action cannot be undone.`)) {
      deleteMutation.mutate(user.id);
    }
  };

  const handleToggleStatus = () => {
    if (user) {
      statusMutation.mutate({ id: user.id, is_active: !user.is_active });
    }
  };

  if (isLoading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
        <CircularProgress />
      </Box>
    );
  }

  if (error || !user) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error" sx={{ mt: 3 }}>
          Failed to load user details
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Button
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/dashboard/users')}
          sx={{ mb: 2 }}
        >
          Back to Users
        </Button>
        <Typography variant="h4" component="h1" gutterBottom>
          User Details
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
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 3, mb: 3 }}>
                <Avatar sx={{ width: 80, height: 80, fontSize: 32 }}>
                  {user.name.charAt(0).toUpperCase()}
                </Avatar>
                <Box sx={{ flexGrow: 1 }}>
                  <Typography variant="h5" gutterBottom>
                    {user.name}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {user.email}
                  </Typography>
                </Box>
                <Chip
                  label={user.is_active ? 'Active' : 'Inactive'}
                  color={user.is_active ? 'success' : 'default'}
                />
              </Box>

              <Divider sx={{ my: 2 }} />

              <Grid container spacing={2}>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <PersonIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        Role
                      </Typography>
                      <Typography variant="body1">
                        {user.role_name}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <SchoolIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        School
                      </Typography>
                      <Typography variant="body1">
                        {user.school_name || '-'}
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
                        {user.email}
                      </Typography>
                    </Box>
                  </Box>
                </Grid>
                <Grid size={{ xs: 12, sm: 6 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                    <BadgeIcon sx={{ mr: 1, color: 'text.secondary' }} />
                    <Box>
                      <Typography variant="body2" color="text.secondary">
                        School ID
                      </Typography>
                      <Typography variant="body1">
                        {user.school_id || '-'}
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
                  onClick={() => navigate(`/dashboard/users/${user.id}`)}
                >
                  Edit User
                </Button>
                <Button
                  variant="outlined"
                  onClick={handleToggleStatus}
                  disabled={statusMutation.isPending}
                >
                  {user.is_active ? 'Deactivate' : 'Activate'}
                </Button>
                <Button
                  variant="outlined"
                  color="error"
                  startIcon={<DeleteIcon />}
                  onClick={handleDeleteClick}
                  disabled={deleteMutation.isPending}
                >
                  Delete User
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
                  {new Date(user.created_at).toLocaleString()}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Updated At
                </Typography>
                <Typography variant="body1">
                  {new Date(user.updated_at).toLocaleString()}
                </Typography>
              </Box>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
};

export default UserDetail;
