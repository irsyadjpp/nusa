import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Stack } from '@mui/material';
import { EvidenceWorkspaceHeader } from '@/features/evidence';
import { EvidenceList, EvidenceFilters, EvidenceUploadPanel, EvidenceReviewPanel } from '@/features/evidence';
import { useEvidences } from '@/services/queries/EvidenceQueryService';
import { useCreateEvidence, useDeleteEvidence } from '@/services/commands/EvidenceCommandService';
import { Evidence } from '@/shared/types/domain';
import { enqueueSnackbar } from 'notistack';

const EvidencePage: React.FC = () => {
  const [selectedEvidence, setSelectedEvidence] = useState<Evidence | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  const [isReviewing, setIsReviewing] = useState(false);
  const [filters, setFilters] = useState<{ student_id?: string; assessment_id?: string; status?: string }>({});
  const [currentUserId] = useState('user-1'); // TODO: Get from auth context

  // Query Evidences
  const { data: evidences = [], isLoading, refetch } = useEvidences(filters);

  // Mutations
  const createEvidenceMutation = useCreateEvidence({
    onSuccess: () => {
      enqueueSnackbar('Evidence uploaded successfully', { variant: 'success' });
      setIsUploading(false);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to upload Evidence', { variant: 'error' });
    },
  });

  const deleteEvidenceMutation = useDeleteEvidence({
    onSuccess: () => {
      enqueueSnackbar('Evidence deleted successfully', { variant: 'success' });
      setSelectedEvidence(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete Evidence', { variant: 'error' });
    },
  });

  const handleUploadNew = () => {
    setIsUploading(true);
    setSelectedEvidence(null);
    setIsReviewing(false);
  };

  const handleSelectEvidence = (evidence: Evidence) => {
    setSelectedEvidence(evidence);
    setIsUploading(false);
    setIsReviewing(false);
  };

  const handleReviewEvidence = () => {
    setIsReviewing(true);
    setIsUploading(false);
  };

  const handleDeleteEvidence = () => {
    if (selectedEvidence && window.confirm('Are you sure you want to delete this Evidence?')) {
      deleteEvidenceMutation.mutate(selectedEvidence.id);
    }
  };

  const handleUploadEvidence = (values: any) => {
    createEvidenceMutation.mutate({ data: values, userId: currentUserId });
  };

  const handleFilterChange = (newFilters: any) => {
    setFilters(newFilters);
  };

  const handleClearFilters = () => {
    setFilters({});
  };

  return (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <EvidenceWorkspaceHeader
        title="Evidence Workspace"
        onUploadNew={handleUploadNew}
        onRefresh={() => refetch()}
        showUploadButton
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - Evidence List and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <EvidenceFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <EvidenceList
                evidences={evidences}
                selectedId={selectedEvidence?.id}
                onSelect={handleSelectEvidence}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - Evidence Detail/Upload/Review */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isUploading && (
              <EvidenceUploadPanel
                onUpload={handleUploadEvidence}
                onCancel={() => setIsUploading(false)}
              />
            )}

            {isReviewing && selectedEvidence && (
              <EvidenceReviewPanel
                evidence={selectedEvidence}
                onApprove={() => setIsReviewing(false)}
                onReject={() => setIsReviewing(false)}
              />
            )}

            {selectedEvidence && !isUploading && !isReviewing && (
              <Box>
                <EvidenceReviewPanel
                  evidence={selectedEvidence}
                  onApprove={() => setIsReviewing(false)}
                  onReject={() => setIsReviewing(false)}
                />
                <Divider sx={{ my: 2 }} />
                <Stack direction="row" spacing={2} justifyContent="flex-end">
                  <Button variant="contained" onClick={handleReviewEvidence}>
                    Review
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteEvidence}>
                    Delete
                  </Button>
                </Stack>
              </Box>
            )}

            {!selectedEvidence && !isUploading && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select an Evidence from the list or upload new evidence
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default EvidencePage;
