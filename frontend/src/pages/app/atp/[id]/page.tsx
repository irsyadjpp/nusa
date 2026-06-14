import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Box, Button, Container, Paper, Typography, Breadcrumbs, Link, CircularProgress, Alert, Stack } from '@mui/material';
import { useATP, useATPSet } from '@/services/queries/ATPQueryService';
import { useApproveATPSet } from '@/services/commands/ATPCommandService';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import NavigateNextIcon from '@mui/icons-material/NavigateNext';

const ATPDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const { data: atp, isLoading, error } = useATP(id || '');
  const { data: atpSet } = useATPSet(atp?.atp_set_id || '');
  
  const approveATPSetMutation = useApproveATPSet({
    onSuccess: () => {
      navigate('/atp');
    },
  });

  const handleApproveATPSet = () => {
    if (atpSet && window.confirm('Are you sure you want to approve this ATP Set?')) {
      approveATPSetMutation.mutate(atpSet.id);
    }
  };

  const handleCreateModulAjar = () => {
    navigate('/modul-ajar/create', { state: { selectedATP: atp } });
  };

  const handleViewAchievement = () => {
    navigate('/achievement');
  };

  const handleBack = () => {
    navigate('/atp');
  };

  if (isLoading) {
    return (
      <Container maxWidth="xl">
        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '50vh' }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error">
          Failed to load ATP details. Please try again.
        </Alert>
      </Container>
    );
  }

  if (!atp) {
    return (
      <Container maxWidth="xl">
        <Alert severity="info">
          ATP not found.
        </Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl">
      <Box sx={{ py: 3 }}>
        <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
          <Link underline="hover" color="inherit" onClick={handleBack}>
            ATP
          </Link>
          <Typography color="text.primary">{atp.id}</Typography>
        </Breadcrumbs>

        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
          <Typography variant="h4" component="h1">
            ATP Detail
          </Typography>
          <Stack direction="row" spacing={2}>
            <Button
              variant="outlined"
              startIcon={<ArrowBackIcon />}
              onClick={handleBack}
            >
              Back to ATP List
            </Button>
            <Button
              variant="contained"
              startIcon={<NavigateNextIcon />}
              onClick={handleCreateModulAjar}
            >
              Create Modul Ajar from this ATP
            </Button>
            <Button
              variant="outlined"
              onClick={handleViewAchievement}
            >
              View Achievement
            </Button>
          </Stack>
        </Box>

        <Paper sx={{ p: 3, mb: 3 }}>
          <Typography variant="h6" gutterBottom>
            ATP Information
          </Typography>
          <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
            <Box>
              <Typography variant="body2" color="text.secondary">
                ATP ID
              </Typography>
              <Typography variant="body1">
                {atp.id}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                ATP Set ID
              </Typography>
              <Typography variant="body1">
                {atp.atp_set_id}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                TP ID
              </Typography>
              <Typography variant="body1">
                {atp.tp_id}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Sequence Number
              </Typography>
              <Typography variant="body1">
                {atp.sequence_number}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Week Number
              </Typography>
              <Typography variant="body1">
                {atp.week_number}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Estimated Hours
              </Typography>
              <Typography variant="body1">
                {atp.estimated_hours}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Status
              </Typography>
              <Typography variant="body1">
                {atp.status}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Created
              </Typography>
              <Typography variant="body1">
                {new Date(atp.created_at).toLocaleString()}
              </Typography>
            </Box>
            <Box>
              <Typography variant="body2" color="text.secondary">
                Updated
              </Typography>
              <Typography variant="body1">
                {new Date(atp.updated_at).toLocaleString()}
              </Typography>
            </Box>
          </Box>
        </Paper>

        {atpSet && (
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              ATP Set Actions
            </Typography>
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
              <Box>
                <Typography variant="body2">
                  ATP Set: {atpSet.id}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Status: {atpSet.status}
                </Typography>
              </Box>
              {atpSet.status !== 'APPROVED' && (
                <Button
                  variant="contained"
                  color="success"
                  onClick={handleApproveATPSet}
                  disabled={approveATPSetMutation.isPending}
                >
                  {approveATPSetMutation.isPending ? 'Approving...' : 'Approve ATP Set'}
                </Button>
              )}
            </Box>
          </Paper>
        )}
      </Box>
    </Container>
  );
};

export default ATPDetailPage;