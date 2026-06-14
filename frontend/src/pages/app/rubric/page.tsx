import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Stack } from '@mui/material';
import { RubricDesignerHeader } from '@/features/rubric';
import { RubricList, RubricFilters, RubricForm, RubricPreviewPanel } from '@/features/rubric';
import { useRubrics } from '@/services/queries/RubricQueryService';
import { useCreateRubric, useUpdateRubric, useDeleteRubric } from '@/services/commands/RubricCommandService';
import { Rubric } from '@/shared/types/domain';
import { enqueueSnackbar } from 'notistack';

const RubricPage: React.FC = () => {
  const [selectedRubric, setSelectedRubric] = useState<Rubric | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [showPreview, setShowPreview] = useState(false);
  const [filters, setFilters] = useState<{ assessment_id?: string; status?: string }>({});
  const [currentUserId] = useState('user-1'); // TODO: Get from auth context

  // Query Rubrics
  const { data: rubrics = [], isLoading, refetch } = useRubrics(filters);

  // Mutations
  const createRubricMutation = useCreateRubric({
    onSuccess: () => {
      enqueueSnackbar('Rubric created successfully', { variant: 'success' });
      setIsCreating(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to create Rubric', { variant: 'error' });
    },
  });

  const updateRubricMutation = useUpdateRubric({
    onSuccess: () => {
      enqueueSnackbar('Rubric updated successfully', { variant: 'success' });
      setIsEditing(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to update Rubric', { variant: 'error' });
    },
  });

  const deleteRubricMutation = useDeleteRubric({
    onSuccess: () => {
      enqueueSnackbar('Rubric deleted successfully', { variant: 'success' });
      setSelectedRubric(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete Rubric', { variant: 'error' });
    },
  });

  const handleCreateNew = () => {
    setIsCreating(true);
    setSelectedRubric(null);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleSelectRubric = (rubric: Rubric) => {
    setSelectedRubric(rubric);
    setIsCreating(false);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleEditRubric = () => {
    setIsEditing(true);
    setShowPreview(false);
  };

  const handleDeleteRubric = () => {
    if (selectedRubric && window.confirm('Are you sure you want to delete this Rubric?')) {
      deleteRubricMutation.mutate(selectedRubric.id);
    }
  };

  const handlePreview = () => {
    setShowPreview(true);
    setIsEditing(false);
  };

  const handleCreateRubric = (values: any) => {
    createRubricMutation.mutate({ data: values, userId: currentUserId });
  };

  const handleUpdateRubric = (values: any) => {
    if (selectedRubric) {
      updateRubricMutation.mutate({ id: selectedRubric.id, data: values });
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
      <RubricDesignerHeader
        title="Rubric Designer"
        onCreateNew={handleCreateNew}
        onRefresh={() => refetch()}
        showCreateButton
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - Rubric List and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <RubricFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <RubricList
                rubrics={rubrics}
                selectedId={selectedRubric?.id}
                onSelect={handleSelectRubric}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - Rubric Detail/Form */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isCreating && (
              <RubricForm
                onSubmit={handleCreateRubric}
                onCancel={() => setIsCreating(false)}
                isEdit={false}
              />
            )}

            {isEditing && selectedRubric && (
              <RubricForm
                initialValues={selectedRubric}
                onSubmit={handleUpdateRubric}
                onCancel={() => setIsEditing(false)}
                isEdit={true}
              />
            )}

            {showPreview && selectedRubric && (
              <Box>
                <RubricPreviewPanel rubric={selectedRubric} />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowPreview(false)}>
                    Close Preview
                  </Button>
                </Box>
              </Box>
            )}

            {selectedRubric && !isCreating && !isEditing && !showPreview && (
              <Box>
                <RubricPreviewPanel rubric={selectedRubric} />
                <Divider sx={{ my: 2 }} />
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button variant="contained" onClick={handleEditRubric}>
                    Edit
                  </Button>
                  <Button variant="outlined" onClick={handlePreview}>
                    Preview
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteRubric}>
                    Delete
                  </Button>
                </Stack>
              </Box>
            )}

            {!selectedRubric && !isCreating && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select a Rubric from the list or create a new one
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default RubricPage;
