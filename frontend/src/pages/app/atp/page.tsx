import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Stack } from '@mui/material';
import { ATPWorkspaceHeader } from '@/features/atp';
import { ATPList, ATPFilters, ATPForm, ATPDetailPanel } from '@/features/atp';
import { useATPs } from '@/services/queries/ATPQueryService';
import { useCreateATP, useUpdateATP, useDeleteATP } from '@/services/commands/ATPCommandService';
import { ATP } from '@/api/atp';
import { enqueueSnackbar } from 'notistack';

const AnnualTeachingPlanPage: React.FC = () => {
  const [selectedATP, setSelectedATP] = useState<ATP | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [filters, setFilters] = useState<{ tp_id?: string; status?: string }>({});

  // Query ATPs
  const { data: atps = [], isLoading, refetch } = useATPs(filters);

  // Mutations
  const createATPMutation = useCreateATP({
    onSuccess: () => {
      enqueueSnackbar('ATP created successfully', { variant: 'success' });
      setIsCreating(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to create ATP', { variant: 'error' });
    },
  });

  const updateATPMutation = useUpdateATP({
    onSuccess: () => {
      enqueueSnackbar('ATP updated successfully', { variant: 'success' });
      setIsEditing(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to update ATP', { variant: 'error' });
    },
  });

  const deleteATPMutation = useDeleteATP({
    onSuccess: () => {
      enqueueSnackbar('ATP deleted successfully', { variant: 'success' });
      setSelectedATP(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete ATP', { variant: 'error' });
    },
  });

  const handleCreateNew = () => {
    setIsCreating(true);
    setSelectedATP(null);
    setIsEditing(false);
  };

  const handleSelectATP = (atp: ATP) => {
    setSelectedATP(atp);
    setIsCreating(false);
    setIsEditing(false);
  };

  const handleEditATP = () => {
    setIsEditing(true);
  };

  const handleDeleteATP = () => {
    if (selectedATP && window.confirm('Are you sure you want to delete this ATP?')) {
      deleteATPMutation.mutate(selectedATP.id);
    }
  };

  const handleCreateATP = (values: any) => {
    createATPMutation.mutate(values);
  };

  const handleUpdateATP = (values: any) => {
    if (selectedATP) {
      updateATPMutation.mutate({ id: selectedATP.id, data: values });
    }
  };

  const handleFilterChange = (newFilters: any) => {
    setFilters(newFilters);
  };

  const handleClearFilters = () => {
    setFilters({});
  };

  return (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <ATPWorkspaceHeader
        title="Annual Teaching Plan (ATP)"
        onCreateNew={handleCreateNew}
        onRefresh={() => refetch()}
        showCreateButton
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - ATP List and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <ATPFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <ATPList
                atps={atps}
                selectedId={selectedATP?.id}
                onSelect={handleSelectATP}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - ATP Detail/Form */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isCreating && (
              <ATPForm
                onSubmit={handleCreateATP}
                onCancel={() => setIsCreating(false)}
                isEdit={false}
              />
            )}

            {isEditing && selectedATP && (
              <ATPForm
                initialValues={selectedATP}
                onSubmit={handleUpdateATP}
                onCancel={() => setIsEditing(false)}
                isEdit={true}
              />
            )}

            {selectedATP && !isCreating && !isEditing && (
              <Box>
                <ATPDetailPanel atp={selectedATP} />
                <Divider sx={{ my: 2 }} />
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button variant="contained" onClick={handleEditATP}>
                    Edit
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteATP}>
                    Delete
                  </Button>
                </Stack>
              </Box>
            )}

            {!selectedATP && !isCreating && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select an ATP from the list or create a new one
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default AnnualTeachingPlanPage;
