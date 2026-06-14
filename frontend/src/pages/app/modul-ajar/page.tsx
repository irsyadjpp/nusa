import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Stack } from '@mui/material';
import { ModulAjarWorkspaceHeader } from '@/features/modul-ajar';
import { ModulAjarList, ModulAjarFilters, ModulAjarForm, ModulAjarPreviewPanel } from '@/features/modul-ajar';
import { useModulAjars } from '@/services/queries/ModulAjarQueryService';
import { useCreateModulAjar, useUpdateModulAjar, useDeleteModulAjar } from '@/services/commands/ModulAjarCommandService';
import { ModulAjar } from '@/shared/types/domain';
import { enqueueSnackbar } from 'notistack';
import { useNavigate } from 'react-router-dom';
import NavigateNextIcon from '@mui/icons-material/NavigateNext';

const ModulAjarPage: React.FC = () => {
  const navigate = useNavigate();
  const [selectedModulAjar, setSelectedModulAjar] = useState<ModulAjar | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [showPreview, setShowPreview] = useState(false);
  const [filters, setFilters] = useState<{ atp_id?: string; tp_id?: string; status?: string }>({});

  // Query Modul Ajars
  const { data: modulAjars = [], isLoading, refetch } = useModulAjars(filters);

  // Mutations
  const createModulAjarMutation = useCreateModulAjar({
    onSuccess: () => {
      enqueueSnackbar('Modul Ajar created successfully', { variant: 'success' });
      setIsCreating(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to create Modul Ajar', { variant: 'error' });
    },
  });

  const updateModulAjarMutation = useUpdateModulAjar({
    onSuccess: () => {
      enqueueSnackbar('Modul Ajar updated successfully', { variant: 'success' });
      setIsEditing(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to update Modul Ajar', { variant: 'error' });
    },
  });

  const deleteModulAjarMutation = useDeleteModulAjar({
    onSuccess: () => {
      enqueueSnackbar('Modul Ajar deleted successfully', { variant: 'success' });
      setSelectedModulAjar(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete Modul Ajar', { variant: 'error' });
    },
  });

  const handleCreateNew = () => {
    setIsCreating(true);
    setSelectedModulAjar(null);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleSelectModulAjar = (modulAjar: ModulAjar) => {
    setSelectedModulAjar(modulAjar);
    setIsCreating(false);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleEditModulAjar = () => {
    setIsEditing(true);
    setShowPreview(false);
  };

  const handleDeleteModulAjar = () => {
    if (selectedModulAjar && window.confirm('Are you sure you want to delete this Modul Ajar?')) {
      deleteModulAjarMutation.mutate(selectedModulAjar.id);
    }
  };

  const handlePreview = () => {
    setShowPreview(true);
    setIsEditing(false);
  };

  const handleCreateModulAjar = (values: any) => {
    createModulAjarMutation.mutate(values);
  };

  const handleUpdateModulAjar = (values: any) => {
    if (selectedModulAjar) {
      updateModulAjarMutation.mutate({ id: selectedModulAjar.id, data: values });
    }
  };

  const handleFilterChange = (newFilters: any) => {
    setFilters(newFilters);
  };

  const handleClearFilters = () => {
    setFilters({});
  };

  const handleCreateAssessment = () => {
    if (selectedModulAjar) {
      navigate('/assessment/create', { state: { modulAjarId: selectedModulAjar.id } });
    }
  };

  return (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <ModulAjarWorkspaceHeader
        title="Modul Ajar"
        onCreateNew={handleCreateNew}
        onRefresh={() => refetch()}
        showCreateButton
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - Modul Ajar List and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <ModulAjarFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <ModulAjarList
                modulAjars={modulAjars}
                selectedId={selectedModulAjar?.id}
                onSelect={handleSelectModulAjar}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - Modul Ajar Detail/Form */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isCreating && (
              <ModulAjarForm
                onSubmit={handleCreateModulAjar}
                onCancel={() => setIsCreating(false)}
                isEdit={false}
              />
            )}

            {isEditing && selectedModulAjar && (
              <ModulAjarForm
                initialValues={selectedModulAjar}
                onSubmit={handleUpdateModulAjar}
                onCancel={() => setIsEditing(false)}
                isEdit={true}
              />
            )}

            {showPreview && selectedModulAjar && (
              <Box>
                <ModulAjarPreviewPanel modulAjar={selectedModulAjar} />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowPreview(false)}>
                    Close Preview
                  </Button>
                </Box>
              </Box>
            )}

            {selectedModulAjar && !isCreating && !isEditing && !showPreview && (
              <Box>
                <ModulAjarPreviewPanel modulAjar={selectedModulAjar} />
                <Divider sx={{ my: 2 }} />
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button variant="contained" onClick={handleEditModulAjar}>
                    Edit
                  </Button>
                  <Button variant="outlined" onClick={handlePreview}>
                    Preview
                  </Button>
                  <Button
                    variant="outlined"
                    startIcon={<NavigateNextIcon />}
                    onClick={handleCreateAssessment}
                  >
                    Create Assessment
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteModulAjar}>
                    Delete
                  </Button>
                </Stack>
              </Box>
            )}

            {!selectedModulAjar && !isCreating && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select a Modul Ajar from the list or create a new one
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default ModulAjarPage;
