import React from 'react';
import { useParams } from 'react-router-dom';
import {
  Box,
  Typography,
  Button,
  Paper,
  CircularProgress,
  Alert,
  Container,
  Breadcrumbs,
  Link,
  Chip,
  Grid,
  Divider,
  IconButton,
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import EditIcon from '@mui/icons-material/Edit';
import PublishIcon from '@mui/icons-material/Publish';
import { useNavigate } from 'react-router-dom';
import { useNarrativeReport, useAchievementSummary } from '@/services/queries/NarrativeReportQueryService';
import { publishNarrativeReport } from '@/api/narrative-report';

const NarrativeReportDetailPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const {
    data: report,
    isLoading: isLoadingReport,
    error: reportError,
    refetch: refetchReport,
  } = useNarrativeReport(id!);

  const {
    data: achievementSummary,
    isLoading: isLoadingSummary,
    error: summaryError,
  } = useAchievementSummary(id!);

  const handlePublish = async () => {
    try {
      await publishNarrativeReport(id!);
      refetchReport();
    } catch (error) {
      console.error('Failed to publish narrative report:', error);
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'PUBLISHED':
        return 'success';
      case 'DRAFT':
        return 'default';
      case 'PENDING_REVIEW':
        return 'warning';
      default:
        return 'default';
    }
  };

  if (isLoadingReport || isLoadingSummary) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (reportError || summaryError) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Alert severity="error">{reportError?.message || summaryError?.message || 'Failed to load narrative report'}</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl" sx={{ py: 4 }}>
      <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
        <Link underline="hover" onClick={() => navigate('/narrative-reports')}>
          Narrative Reports
        </Link>
        <Typography color="text.primary">Report Detail</Typography>
      </Breadcrumbs>

      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" component="h1">
          Narrative Report Detail
        </Typography>
        <Button
          variant="outlined"
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/narrative-reports')}
        >
          Back to Reports
        </Button>
      </Box>

      {report && (
        <Grid container spacing={3}>
          <Grid size={{ xs: 12 }}>
            <Paper sx={{ p: 3 }}>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
                <Typography variant="h6">Report Information</Typography>
                <Box sx={{ display: 'flex', gap: 1 }}>
                  {report.status === 'DRAFT' && (
                    <IconButton onClick={handlePublish} title="Publish">
                      <PublishIcon />
                    </IconButton>
                  )}
                  <IconButton
                    onClick={() => navigate(`/narrative-reports/${id}/edit`)}
                    title="Edit"
                  >
                    <EditIcon />
                  </IconButton>
                </Box>
              </Box>
              <Grid container spacing={2}>
                <Grid size={{ xs: 12, md: 6 }}>
                  <Typography variant="body2" color="text.secondary">
                    Report ID
                  </Typography>
                  <Typography variant="body1">{report.id}</Typography>
                </Grid>
                <Grid size={{ xs: 12, md: 6 }}>
                  <Typography variant="body2" color="text.secondary">
                    Student ID
                  </Typography>
                  <Typography variant="body1">{report.student_id}</Typography>
                </Grid>
                <Grid size={{ xs: 12, md: 6 }}>
                  <Typography variant="body2" color="text.secondary">
                    Period ID
                  </Typography>
                  <Typography variant="body1">{report.period_id}</Typography>
                </Grid>
                <Grid size={{ xs: 12, md: 6 }}>
                  <Typography variant="body2" color="text.secondary">
                    Status
                  </Typography>
                  <Chip
                    label={report.status}
                    color={getStatusColor(report.status) as any}
                    size="small"
                  />
                </Grid>
                <Grid size={{ xs: 12, md: 6 }}>
                  <Typography variant="body2" color="text.secondary">
                    Created By
                  </Typography>
                  <Typography variant="body1">{report.created_by}</Typography>
                </Grid>
                <Grid size={{ xs: 12, md: 6 }}>
                  <Typography variant="body2" color="text.secondary">
                    Created At
                  </Typography>
                  <Typography variant="body1">
                    {new Date(report.created_at).toLocaleString()}
                  </Typography>
                </Grid>
                {report.published_by && (
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Typography variant="body2" color="text.secondary">
                      Published By
                    </Typography>
                    <Typography variant="body1">{report.published_by}</Typography>
                  </Grid>
                )}
                {report.published_at && (
                  <Grid size={{ xs: 12, md: 6 }}>
                    <Typography variant="body2" color="text.secondary">
                      Published At
                    </Typography>
                    <Typography variant="body1">
                      {new Date(report.published_at).toLocaleString()}
                    </Typography>
                  </Grid>
                )}
              </Grid>
            </Paper>
          </Grid>

          <Grid size={{ xs: 12 }}>
            <Paper sx={{ p: 3 }}>
              <Typography variant="h6" gutterBottom>
                Narrative Content
              </Typography>
              <Typography variant="body1" sx={{ whiteSpace: 'pre-wrap' }}>
                {typeof report.narrative_content === 'string' 
                  ? report.narrative_content 
                  : JSON.stringify(report.narrative_content, null, 2)}
              </Typography>
            </Paper>
          </Grid>

          {achievementSummary && (
            <Grid size={{ xs: 12 }}>
              <Paper sx={{ p: 3 }}>
                <Typography variant="h6" gutterBottom>
                  Achievement Summary
                </Typography>
                <Divider sx={{ mb: 2 }} />
                <Typography variant="body2" color="text.secondary">
                  Total Achievements
                </Typography>
                <Typography variant="h4" gutterBottom>
                  {achievementSummary.total_achievements || 0}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Average Mastery
                </Typography>
                <Typography variant="h4" gutterBottom>
                  {achievementSummary.average_mastery || 0}%
                </Typography>
                {achievementSummary.achievements && achievementSummary.achievements.length > 0 && (
                  <Box>
                    <Typography variant="body2" color="text.secondary" gutterBottom>
                      Individual Achievements
                    </Typography>
                    {achievementSummary.achievements.map((achievement: any) => (
                      <Box key={achievement.competency_id} sx={{ mb: 2 }}>
                        <Typography variant="subtitle1">{achievement.competency_name}</Typography>
                        <Typography variant="body2" color="text.secondary">
                          Mastery Level: {achievement.mastery_level}
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          Score: {achievement.score} / {achievement.max_score}
                        </Typography>
                      </Box>
                    ))}
                  </Box>
                )}
              </Paper>
            </Grid>
          )}

          <Grid size={{ xs: 12 }}>
            <Paper sx={{ p: 3 }}>
              <Typography variant="h6" gutterBottom>
                Actions
              </Typography>
              <Box sx={{ display: 'flex', gap: 2 }}>
                <Button
                  variant="contained"
                  onClick={() => navigate(`/narrative-reports/${id}/edit`)}
                >
                  Edit Report
                </Button>
                {report.status === 'DRAFT' && (
                  <Button variant="outlined" onClick={handlePublish}>
                    Publish Report
                  </Button>
                )}
                <Button
                  variant="outlined"
                  onClick={() => navigate('/narrative-reports')}
                >
                  Back to Reports
                </Button>
              </Box>
            </Paper>
          </Grid>
        </Grid>
      )}
    </Container>
  );
};

export default NarrativeReportDetailPage;
