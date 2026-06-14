import React, { useState } from 'react';
import { Box, Grid, Paper, Divider, Button, Typography } from '@mui/material';
import { NarrativeReportBuilderHeader, ReportList, ReportFilters, ReportBuilderPanel, ReportPreview, ReportActions } from '@/features/narrative-report';
import { useNarrativeReports } from '@/services/queries/NarrativeReportQueryService';
import { useDeleteNarrativeReport, usePublishNarrativeReport } from '@/services/commands/NarrativeReportCommandService';
import { NarrativeReport } from '@/shared/types/domain';
import { enqueueSnackbar } from 'notistack';

const NarrativeReportPage: React.FC = () => {
  const [selectedReport, setSelectedReport] = useState<NarrativeReport | null>(null);
  const [isCreating, setIsCreating] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [showPreview, setShowPreview] = useState(false);
  const [filters, setFilters] = useState<{ student_id?: string; period_id?: string; status?: string }>({});

  // Query Reports
  const { data: reports = [], isLoading, refetch } = useNarrativeReports(filters);

  // Mutations
  const deleteReportMutation = useDeleteNarrativeReport({
    onSuccess: () => {
      enqueueSnackbar('Narrative Report deleted successfully', { variant: 'success' });
      setSelectedReport(null);
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to delete Narrative Report', { variant: 'error' });
    },
  });

  const publishReportMutation = usePublishNarrativeReport({
    onSuccess: () => {
      enqueueSnackbar('Narrative Report published successfully', { variant: 'success' });
      refetch();
    },
    onError: (error: any) => {
      enqueueSnackbar(error.message || 'Failed to publish Narrative Report', { variant: 'error' });
    },
  });

  const handleCreateNew = () => {
    setIsCreating(true);
    setSelectedReport(null);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleSelectReport = (report: NarrativeReport) => {
    setSelectedReport(report);
    setIsCreating(false);
    setIsEditing(false);
    setShowPreview(false);
  };

  const handleEditReport = () => {
    setIsEditing(true);
    setShowPreview(false);
  };

  const handleDeleteReport = () => {
    if (selectedReport && window.confirm('Are you sure you want to delete this Narrative Report?')) {
      deleteReportMutation.mutate(selectedReport.id);
    }
  };

  const handlePreview = () => {
    setShowPreview(true);
    setIsEditing(false);
  };

  const handlePublish = () => {
    if (selectedReport) {
      publishReportMutation.mutate(selectedReport.id);
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
      <NarrativeReportBuilderHeader
        title="Narrative Report Builder"
        onCreateNew={handleCreateNew}
        onRefresh={() => refetch()}
        showCreateButton
      />

      <Grid container spacing={2} sx={{ flex: 1, overflow: 'hidden' }}>
        {/* Left Panel - Report List and Filters */}
        <Grid size={{ xs: 12, md: 4 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <Box sx={{ p: 2 }}>
              <ReportFilters
                filters={filters}
                onFilterChange={handleFilterChange}
                onClearFilters={handleClearFilters}
              />
            </Box>
            <Divider />
            <Box sx={{ flex: 1, overflow: 'auto' }}>
              <ReportList
                reports={reports}
                selectedId={selectedReport?.id}
                onSelect={handleSelectReport}
                loading={isLoading}
              />
            </Box>
          </Paper>
        </Grid>

        {/* Right Panel - Report Detail/Builder */}
        <Grid size={{ xs: 12, md: 8 }} sx={{ height: '100%', overflow: 'auto' }}>
          <Paper sx={{ height: '100%', p: 2 }}>
            {isCreating && (
              <Box>
                <Typography variant="h6" gutterBottom>
                  Create New Narrative Report
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Use the Report Wizard to create a new narrative report
                </Typography>
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setIsCreating(false)}>
                    Cancel
                  </Button>
                </Box>
              </Box>
            )}

            {isEditing && selectedReport && (
              <ReportBuilderPanel
                report={selectedReport}
                onEdit={() => {}}
              />
            )}

            {showPreview && selectedReport && (
              <Box>
                <ReportPreview report={selectedReport} />
                <Box sx={{ mt: 2 }}>
                  <Button variant="outlined" onClick={() => setShowPreview(false)}>
                    Close Preview
                  </Button>
                </Box>
              </Box>
            )}

            {selectedReport && !isCreating && !isEditing && !showPreview && (
              <Box>
                <ReportPreview report={selectedReport} />
                <Divider sx={{ my: 2 }} />
                <Box sx={{ display: 'flex', gap: 1, justifyContent: 'flex-end' }}>
                  <Button variant="contained" onClick={handleEditReport}>
                    Edit
                  </Button>
                  <Button variant="outlined" onClick={handlePreview}>
                    Preview
                  </Button>
                  <Button variant="outlined" color="success" onClick={handlePublish}>
                    Publish
                  </Button>
                  <Button variant="outlined" color="error" onClick={handleDeleteReport}>
                    Delete
                  </Button>
                </Box>
                <Divider sx={{ my: 2 }} />
                <ReportActions
                  reportId={selectedReport.id}
                  onRefreshAchievement={() => refetch()}
                />
              </Box>
            )}

            {!selectedReport && !isCreating && (
              <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                Select a Narrative Report from the list or create a new one
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default NarrativeReportPage;
