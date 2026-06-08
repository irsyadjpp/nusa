import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Stack } from '@mui/material';
import { EvaluationWorkspaceHeader } from '@/features/evaluation';
import { EvaluationQueue, EvaluationFilters, EvaluationPanel, EvaluationForm, EvaluationPreview } from '@/features/evaluation';
import { useEvaluations } from '@/services/queries/EvaluationQueryService';
import { useCreateEvaluation, useUpdateEvaluation, useDeleteEvaluation } from '@/services/commands/EvaluationCommandService';
import { Evaluation } from '@/api/evaluation';
import { enqueueSnackbar } from 'notistack';

const EvaluationPage: React.FC = () => {
  const [selectedEvaluation, setSelectedEvaluation] = useState<Evaluation | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [showPreview, setShowPreview] = useState(false);
  const [filters, setFilters] = useState<{ student_id?: string; rubric_id?: string; evidence_id?: string }>({});
  const [currentUserId] = useState('user-1'); // TODO: Get from auth context

  // Query Evaluations
  const { data: evaluations = [], isLoading, refetch } = useEvaluations(filters);

  // Mutations
  const createEvaluationMutation = useCreateEvaluation({
    onSuccess: () => {
      enqueueSnackbar('Evaluation created successfully', { variant: 'success' });
      setIsCreating(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to create Evaluation', { variant: 'error' });
    },
  });

  const updateEvaluationMutation = useUpdateEvaluation({
    onSuccess: () => {
      enqueueSnackbar('Evaluation updated successfully', { variant: 'success' });
      setIsEditing(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to update Evaluation', { variant: 'error' });
    },
  });

  const deleteEvaluationMutation = useDeleteEvaluation({
    onSuccess: () => {
      enqueueSnackbar('Evaluation deleted successfully', { variant: 'success' });
      setSelectedEvaluation(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete Evaluation', { variant: 'error' });
    },
  });

  const handleCreateNew = () => {
    setIsCreating(true);
    setSelectedEvaluation(null);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleSelectEvaluation = (evaluation: Evaluation) => {
    setSelectedEvaluation(evaluation);
    setIsCreating(false);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleEditEvaluation = () => {
    setIsEditing(true);
    setShowPreview(false);
  };

  const handleDeleteEvaluation = () => {
    if (selectedEvaluation && window.confirm('Are you sure you want to delete this Evaluation?')) {
      deleteEvaluationMutation.mutate(selectedEvaluation.id);
    }
  };

  const handlePreview = () => {
    setShowPreview(true);
    setIsEditing(false);
  };

  const handleCreateEvaluation = (values: any) => {
    createEvaluationMutation.mutate({ data: values, userId: currentUserId });
  };

  const handleUpdateEvaluation = (values: any) => {
    if (selectedEvaluation) {
      updateEvaluationMutation.mutate({ id: selectedEvaluation.id, data: values });
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
      <Stack direction="row" justifyContent="space-between" alignItems="center" sx={{ mb: 3 }}>
        <EvaluationWorkspaceHeader
          title="Evaluation Workspace"
          onRefresh={() => refetch()}
        />
        <Button variant="contained" onClick={handleCreateNew}>
          Create New Evaluation
        </Button>
      </Stack>

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - Evaluation Queue and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <EvaluationFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <EvaluationQueue
                evaluations={evaluations}
                selectedId={selectedEvaluation?.id}
                onSelect={handleSelectEvaluation}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - Evaluation Detail/Form */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isCreating && (
              <EvaluationForm
                onSubmit={handleCreateEvaluation}
                onCancel={() => setIsCreating(false)}
                isEdit={false}
              />
            )}

            {isEditing && selectedEvaluation && (
              <EvaluationForm
                initialValues={selectedEvaluation}
                onSubmit={handleUpdateEvaluation}
                onCancel={() => setIsEditing(false)}
                isEdit={true}
              />
            )}

            {showPreview && selectedEvaluation && (
              <Box>
                <EvaluationPreview evaluation={selectedEvaluation} />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowPreview(false)}>
                    Close Preview
                  </Button>
                </Box>
              </Box>
            )}

            {selectedEvaluation && !isCreating && !isEditing && !showPreview && (
              <Box>
                <EvaluationPanel evaluation={selectedEvaluation} />
                <Divider sx={{ my: 2 }} />
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button variant="contained" onClick={handleEditEvaluation}>
                    Edit
                  </Button>
                  <Button variant="outlined" onClick={handlePreview}>
                    Preview
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteEvaluation}>
                    Delete
                  </Button>
                </Stack>
              </Box>
            )}

            {!selectedEvaluation && !isCreating && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select an Evaluation from the queue or create a new one
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default EvaluationPage;
