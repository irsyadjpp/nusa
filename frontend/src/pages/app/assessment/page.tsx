import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Stack } from '@mui/material';
import { AssessmentDesignerHeader } from '@/features/assessment';
import { AssessmentList, AssessmentFilters, AssessmentForm, AssessmentPreviewPanel, AssessmentApprovalPanel } from '@/features/assessment';
import { useAssessments } from '@/services/queries/AssessmentQueryService';
import { useCreateAssessment, useUpdateAssessment, useDeleteAssessment, useApproveAssessment, useRejectAssessment } from '@/services/commands/AssessmentCommandService';
import { Assessment } from '@/api/assessment';
import { enqueueSnackbar } from 'notistack';

const AssessmentPage: React.FC = () => {
  const [selectedAssessment, setSelectedAssessment] = useState<Assessment | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [showPreview, setShowPreview] = useState(false);
  const [showApproval, setShowApproval] = useState(false);
  const [filters, setFilters] = useState<{ tp_id?: string; assessment_type?: string; status?: string }>({});

  // Query Assessments
  const { data: assessments = [], isLoading, refetch } = useAssessments(filters);

  // Mutations
  const createAssessmentMutation = useCreateAssessment({
    onSuccess: () => {
      enqueueSnackbar('Assessment created successfully', { variant: 'success' });
      setIsCreating(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to create Assessment', { variant: 'error' });
    },
  });

  const updateAssessmentMutation = useUpdateAssessment({
    onSuccess: () => {
      enqueueSnackbar('Assessment updated successfully', { variant: 'success' });
      setIsEditing(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to update Assessment', { variant: 'error' });
    },
  });

  const deleteAssessmentMutation = useDeleteAssessment({
    onSuccess: () => {
      enqueueSnackbar('Assessment deleted successfully', { variant: 'success' });
      setSelectedAssessment(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete Assessment', { variant: 'error' });
    },
  });

  const approveAssessmentMutation = useApproveAssessment({
    onSuccess: () => {
      enqueueSnackbar('Assessment approved successfully', { variant: 'success' });
      setShowApproval(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to approve Assessment', { variant: 'error' });
    },
  });

  const rejectAssessmentMutation = useRejectAssessment({
    onSuccess: () => {
      enqueueSnackbar('Assessment rejected', { variant: 'warning' });
      setShowApproval(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to reject Assessment', { variant: 'error' });
    },
  });

  const handleCreateNew = () => {
    setIsCreating(true);
    setSelectedAssessment(null);
    setIsEditing(false);
    setShowPreview(false);
    setShowApproval(false);
  };

  const handleSelectAssessment = (assessment: Assessment) => {
    setSelectedAssessment(assessment);
    setIsCreating(false);
    setIsEditing(false);
    setShowPreview(false);
    setShowApproval(false);
  };

  const handleEditAssessment = () => {
    setIsEditing(true);
    setShowPreview(false);
    setShowApproval(false);
  };

  const handleDeleteAssessment = () => {
    if (selectedAssessment && window.confirm('Are you sure you want to delete this Assessment?')) {
      deleteAssessmentMutation.mutate(selectedAssessment.id);
    }
  };

  const handlePreview = () => {
    setShowPreview(true);
    setIsEditing(false);
    setShowApproval(false);
  };

  const handleApproval = () => {
    setShowApproval(true);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleCreateAssessment = (values: any) => {
    createAssessmentMutation.mutate(values);
  };

  const handleUpdateAssessment = (values: any) => {
    if (selectedAssessment) {
      updateAssessmentMutation.mutate({ id: selectedAssessment.id, data: values });
    }
  };

  const handleFilterChange = (newFilters: any) => {
    setFilters(newFilters);
  };

  const handleClearFilters = () => {
    setFilters({});
  };

  const handleApprove = () => {
    if (selectedAssessment) {
      approveAssessmentMutation.mutate(selectedAssessment.id);
    }
  };

  const handleReject = () => {
    if (selectedAssessment) {
      rejectAssessmentMutation.mutate(selectedAssessment.id);
    }
  };

  return (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <AssessmentDesignerHeader
        title="Assessment Designer"
        onCreateNew={handleCreateNew}
        onRefresh={() => refetch()}
        showCreateButton
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - Assessment List and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <AssessmentFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <AssessmentList
                assessments={assessments}
                selectedId={selectedAssessment?.id}
                onSelect={handleSelectAssessment}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - Assessment Detail/Form */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isCreating && (
              <AssessmentForm
                onSubmit={handleCreateAssessment}
                onCancel={() => setIsCreating(false)}
                isEdit={false}
              />
            )}

            {isEditing && selectedAssessment && (
              <AssessmentForm
                initialValues={selectedAssessment}
                onSubmit={handleUpdateAssessment}
                onCancel={() => setIsEditing(false)}
                isEdit={true}
              />
            )}

            {showPreview && selectedAssessment && (
              <Box>
                <AssessmentPreviewPanel assessment={selectedAssessment} />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowPreview(false)}>
                    Close Preview
                  </Button>
                </Box>
              </Box>
            )}

            {showApproval && selectedAssessment && (
              <Box>
                <AssessmentApprovalPanel
                  onApprove={handleApprove}
                  onReject={handleReject}
                />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowApproval(false)}>
                    Cancel
                  </Button>
                </Box>
              </Box>
            )}

            {selectedAssessment && !isCreating && !isEditing && !showPreview && !showApproval && (
              <Box>
                <AssessmentPreviewPanel assessment={selectedAssessment} />
                <Divider sx={{ my: 2 }} />
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button variant="contained" onClick={handleEditAssessment}>
                    Edit
                  </Button>
                  <Button variant="outlined" onClick={handlePreview}>
                    Preview
                  </Button>
                  <Button variant="outlined" color="success" onClick={handleApproval}>
                    Approve
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteAssessment}>
                    Delete
                  </Button>
                </Stack>
              </Box>
            )}

            {!selectedAssessment && !isCreating && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select an Assessment from the list or create a new one
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default AssessmentPage;
