import React, { useState } from 'react';
import { Box, Grid, Paper, Button, Typography, Container, Breadcrumbs, Link, Divider } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { useModulAjarSets } from '@/services/queries/ModulAjarQueryService';
import { useDeleteModulAjarSet, useApproveModulAjarSet } from '@/services/commands/ModulAjarCommandService';
import { ModulAjarSet } from '@/shared/types/domain';
import { enqueueSnackbar } from 'notistack';

const ModulAjarSetPage: React.FC = () => {
  const navigate = useNavigate();
  const [selectedModulAjarSet, setSelectedModulAjarSet] = useState<ModulAjarSet | null>(null);
  const [filters] = useState<{ atp_set_id?: string; status?: string }>({});

  const { data: modulAjarSets = [], isLoading, refetch } = useModulAjarSets(filters);

  const deleteModulAjarSetMutation = useDeleteModulAjarSet({
    onSuccess: () => {
      enqueueSnackbar('Modul Ajar Set deleted successfully', { variant: 'success' });
      setSelectedModulAjarSet(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete Modul Ajar Set', { variant: 'error' });
    },
  });

  const approveModulAjarSetMutation = useApproveModulAjarSet({
    onSuccess: () => {
      enqueueSnackbar('Modul Ajar Set approved successfully', { variant: 'success' });
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to approve Modul Ajar Set', { variant: 'error' });
    },
  });

  const handleSelectModulAjarSet = (modulAjarSet: ModulAjarSet) => {
    setSelectedModulAjarSet(modulAjarSet);
  };

  const handleDeleteModulAjarSet = () => {
    if (selectedModulAjarSet && window.confirm('Are you sure you want to delete this Modul Ajar Set?')) {
      deleteModulAjarSetMutation.mutate(selectedModulAjarSet.id);
    }
  };

  const handleApproveModulAjarSet = () => {
    if (selectedModulAjarSet && window.confirm('Are you sure you want to approve this Modul Ajar Set?')) {
      approveModulAjarSetMutation.mutate(selectedModulAjarSet.id);
    }
  };

  const handleCreateModulAjarFromSet = () => {
    navigate('/modul-ajar/create', { state: { selectedModulAjarSet } });
  };

  const renderModulAjarSetList = () => {
    if (isLoading) {
      return <Typography>Loading Modul Ajar Sets...</Typography>;
    }

    if (!modulAjarSets || modulAjarSets.length === 0) {
      return <Typography>No Modul Ajar Sets found</Typography>;
    }

    return (
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        {modulAjarSets.map((modulAjarSet: ModulAjarSet) => (
          <Paper
            key={modulAjarSet.id}
            sx={{
              p: 2,
              cursor: 'pointer',
              border: selectedModulAjarSet?.id === modulAjarSet.id ? 2 : 1,
              borderColor: selectedModulAjarSet?.id === modulAjarSet.id ? 'primary.main' : 'divider',
            }}
            onClick={() => handleSelectModulAjarSet(modulAjarSet)}
          >
            <Typography variant="h6">{modulAjarSet.id}</Typography>
            <Typography variant="body2" color="text.secondary">
              ATP Set: {modulAjarSet.atp_set_id}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Status: {modulAjarSet.status}
            </Typography>
            <Typography variant="body2" color="text.secondary">
              Created: {new Date(modulAjarSet.created_at).toLocaleDateString()}
            </Typography>
          </Paper>
        ))}
      </Box>
    );
  };

  return (
    <Container maxWidth="xl">
      <Box sx={{ py: 3 }}>
        <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
          <Link underline="hover" color="inherit" onClick={() => navigate('/learning-design')}>
            Learning Design
          </Link>
          <Typography color="text.primary">Modul Ajar Sets</Typography>
        </Breadcrumbs>

        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
          <Typography variant="h4" component="h1">
            Modul Ajar Sets
          </Typography>
          <Button
            variant="contained"
            onClick={() => navigate('/modul-ajar/create')}
          >
            Create Modul Ajar
          </Button>
        </Box>

        <Grid container spacing={3}>
          {/* Left Panel - Modul Ajar Set List */}
          <Grid size={{ xs: 12, md: selectedModulAjarSet ? 6 : 12 }}>
            <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
              <Box sx={{ p: 2 }}>
                <Typography variant="h6" gutterBottom>
                  Filters
                </Typography>
                <Box sx={{ mt: 2 }}>
                  <Typography variant="body2" color="text.secondary" gutterBottom>
                    Filter by ATP Set or status coming soon
                  </Typography>
                </Box>
              </Box>
              <Divider />
              <Box sx={{ flex: 1, overflow: 'auto', p: 2 }}>
                {renderModulAjarSetList()}
              </Box>
            </Paper>
          </Grid>

          {/* Right Panel - Modul Ajar Set Detail */}
          {selectedModulAjarSet && (
            <Grid size={{ xs: 12, md: 6 }}>
              <Paper sx={{ height: '100%', p: 2 }}>
                <Typography variant="h6" gutterBottom>
                  Modul Ajar Set Detail
                </Typography>
                <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      ID
                    </Typography>
                    <Typography variant="body1">
                      {selectedModulAjarSet.id}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      ATP Set ID
                    </Typography>
                    <Typography variant="body1">
                      {selectedModulAjarSet.atp_set_id}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Version No
                    </Typography>
                    <Typography variant="body1">
                      {selectedModulAjarSet.version_no}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Status
                    </Typography>
                    <Typography variant="body1">
                      {selectedModulAjarSet.status}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Created By
                    </Typography>
                    <Typography variant="body1">
                      {selectedModulAjarSet.generated_by}
                    </Typography>
                  </Box>
                  <Box>
                    <Typography variant="body2" color="text.secondary">
                      Created
                    </Typography>
                    <Typography variant="body1">
                      {new Date(selectedModulAjarSet.created_at).toLocaleString()}
                    </Typography>
                  </Box>
                </Box>
                
                <Divider sx={{ my: 2 }} />
                
                <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                  <Button
                    variant="contained"
                    onClick={handleCreateModulAjarFromSet}
                  >
                    Create Modul Ajar from this Set
                  </Button>
                  {selectedModulAjarSet.status !== 'APPROVED' && (
                    <Button
                      variant="outlined"
                      color="success"
                      onClick={handleApproveModulAjarSet}
                      disabled={approveModulAjarSetMutation.isPending}
                    >
                      {approveModulAjarSetMutation.isPending ? 'Approving...' : 'Approve Modul Ajar Set'}
                    </Button>
                  )}
                  <Button
                    variant="outlined"
                    color="error"
                    onClick={handleDeleteModulAjarSet}
                    disabled={deleteModulAjarSetMutation.isPending}
                  >
                    {deleteModulAjarSetMutation.isPending ? 'Deleting...' : 'Delete Modul Ajar Set'}
                  </Button>
                </Box>
              </Paper>
            </Grid>
          )}
        </Grid>
      </Box>
    </Container>
  );
};

export default ModulAjarSetPage;
