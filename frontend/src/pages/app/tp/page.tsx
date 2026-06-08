import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Stack } from '@mui/material';
import { TPWorkspaceHeader } from '@/features/tp';
import { TPList, TPFilters, TPForm, KKTPCriteriaEditor, TPPreviewPanel, TPApprovalPanel, TPVersionHistory } from '@/features/tp';
import { useTPs } from '@/services/queries/TPQueryService';
import { useCreateTP, useUpdateTP, useDeleteTP } from '@/services/commands/TPCommandService';
import { TP } from '@/api/tp';
import { enqueueSnackbar } from 'notistack';

interface Criteria {
  id: string;
  description: string;
  weight: number;
}

const TeachingPlanPage: React.FC = () => {
  const [selectedTP, setSelectedTP] = useState<TP | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [showPreview, setShowPreview] = useState(false);
  const [showApproval, setShowApproval] = useState(false);
  const [showVersionHistory, setShowVersionHistory] = useState(false);
  const [filters, setFilters] = useState<{ subject_id?: string; phase_id?: string; status?: string }>({});
  const [kktpCriteria, setKKTPCriteria] = useState<Criteria[]>([]);

  // Query TPs
  const { data: tps = [], isLoading, refetch } = useTPs(filters);

  // Mutations
  const createTPMutation = useCreateTP({
    onSuccess: () => {
      enqueueSnackbar('TP created successfully', { variant: 'success' });
      setIsCreating(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to create TP', { variant: 'error' });
    },
  });

  const updateTPMutation = useUpdateTP({
    onSuccess: () => {
      enqueueSnackbar('TP updated successfully', { variant: 'success' });
      setIsEditing(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to update TP', { variant: 'error' });
    },
  });

  const deleteTPMutation = useDeleteTP({
    onSuccess: () => {
      enqueueSnackbar('TP deleted successfully', { variant: 'success' });
      setSelectedTP(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete TP', { variant: 'error' });
    },
  });

  const handleCreateNew = () => {
    setIsCreating(true);
    setSelectedTP(null);
    setIsEditing(false);
    setShowPreview(false);
    setShowApproval(false);
    setShowVersionHistory(false);
    setKKTPCriteria([]);
  };

  const handleSelectTP = (tp: TP) => {
    setSelectedTP(tp);
    setIsCreating(false);
    setIsEditing(false);
    setShowPreview(false);
    setShowApproval(false);
    setShowVersionHistory(false);
    // Initialize KKTP criteria from TP success_criteria if available
    if (tp.success_criteria && typeof tp.success_criteria === 'object') {
      setKKTPCriteria(tp.success_criteria as Criteria[]);
    } else {
      setKKTPCriteria([]);
    }
  };

  const handleEditTP = () => {
    setIsEditing(true);
    setShowPreview(false);
    setShowApproval(false);
    setShowVersionHistory(false);
  };

  const handleDeleteTP = () => {
    if (selectedTP && window.confirm('Are you sure you want to delete this TP?')) {
      deleteTPMutation.mutate(selectedTP.id);
    }
  };

  const handlePreview = () => {
    setShowPreview(true);
    setIsEditing(false);
    setShowApproval(false);
    setShowVersionHistory(false);
  };

  const handleApproval = () => {
    setShowApproval(true);
    setIsEditing(false);
    setShowPreview(false);
    setShowVersionHistory(false);
  };

  const handleVersionHistory = () => {
    setShowVersionHistory(true);
    setIsEditing(false);
    setShowPreview(false);
    setShowApproval(false);
  };

  const handleCreateTP = (values: any) => {
    createTPMutation.mutate(values);
  };

  const handleUpdateTP = (values: any) => {
    if (selectedTP) {
      updateTPMutation.mutate({ id: selectedTP.id, data: values });
    }
  };

  const handleFilterChange = (newFilters: any) => {
    setFilters(newFilters);
  };

  const handleClearFilters = () => {
    setFilters({});
  };

  const handleApproveTP = () => {
    if (selectedTP) {
      enqueueSnackbar('TP approved successfully', { variant: 'success' });
      setShowApproval(false);
      refetch();
    }
  };

  const handleRejectTP = () => {
    if (selectedTP) {
      enqueueSnackbar('TP rejected', { variant: 'warning' });
      setShowApproval(false);
      refetch();
    }
  };

  const handleKKTPCriteriaChange = (criteria: Criteria[]) => {
    setKKTPCriteria(criteria);
  };

  return (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <TPWorkspaceHeader
        title="Teaching Plan (TP)"
        onCreateNew={handleCreateNew}
        onRefresh={() => refetch()}
        showCreateButton
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - TP List and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <TPFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <TPList
                tps={tps}
                selectedId={selectedTP?.id}
                onSelect={handleSelectTP}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - TP Detail/Form */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isCreating && (
              <TPForm
                onSubmit={handleCreateTP}
                onCancel={() => setIsCreating(false)}
                isEdit={false}
              />
            )}

            {isEditing && selectedTP && (
              <TPForm
                initialValues={selectedTP}
                onSubmit={handleUpdateTP}
                onCancel={() => setIsEditing(false)}
                isEdit={true}
              />
            )}

            {showPreview && selectedTP && (
              <Box>
                <TPPreviewPanel tp={selectedTP} />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowPreview(false)}>
                    Close Preview
                  </Button>
                </Box>
              </Box>
            )}

            {showApproval && selectedTP && (
              <Box>
                <TPApprovalPanel
                  onApprove={handleApproveTP}
                  onReject={handleRejectTP}
                />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowApproval(false)}>
                    Cancel
                  </Button>
                </Box>
              </Box>
            )}

            {showVersionHistory && selectedTP && (
              <Box>
                <TPVersionHistory
                  tpSetId={selectedTP.tp_set_id}
                  sequenceNumber={selectedTP.sequence_number}
                />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowVersionHistory(false)}>
                    Close
                  </Button>
                </Box>
              </Box>
            )}

            {selectedTP && !isCreating && !isEditing && !showPreview && !showApproval && !showVersionHistory && (
              <Box>
                <KKTPCriteriaEditor
                  criteria={kktpCriteria}
                  onChange={handleKKTPCriteriaChange}
                />
                <Divider sx={{ my: 2 }} />
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button variant="contained" onClick={handleEditTP}>
                    Edit
                  </Button>
                  <Button variant="outlined" onClick={handlePreview}>
                    Preview
                  </Button>
                  <Button variant="outlined" color="success" onClick={handleApproval}>
                    Approve
                  </Button>
                  <Button variant="outlined" onClick={handleVersionHistory}>
                    Version History
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteTP}>
                    Delete
                  </Button>
                </Stack>
              </Box>
            )}

            {!selectedTP && !isCreating && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select a TP from the list or create a new one
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default TeachingPlanPage;
