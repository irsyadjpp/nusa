import React, { useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import {
  Box,
  Typography,
  Button,
  Paper,
  TextField,
  Alert,
  Container,
  Breadcrumbs,
  Link,
  Grid,
  Divider,
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import {
  createNarrativeReport,
  getAchievementSummary,
} from '@/api/narrative-report';
import { NarrativeContent } from '@/shared/types/domain';

const NarrativeReportGeneratePage: React.FC = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const state = location.state as { studentId?: string; classId?: string };

  const [formData, setFormData] = useState({
    student_id: state?.studentId || '',
    class_id: state?.classId || '',
    subject_id: '',
    period_id: state?.classId || '',
    narrative_content: {} as NarrativeContent,
  });

  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [achievementData, setAchievementData] = useState<any>(null);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleLoadAchievementData = async () => {
    if (!formData.student_id) {
      setError('Please enter a student ID first');
      return;
    }

    try {
      setIsLoading(true);
      setError(null);
      const summary = await getAchievementSummary('', { student_id: formData.student_id });
      setAchievementData(summary);
    } catch (error) {
      setError('Failed to load achievement data');
      console.error('Failed to load achievement data:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      setIsLoading(true);
      setError(null);

      const reportData = {
        student_id: formData.student_id,
        class_id: formData.class_id,
        subject_id: formData.subject_id,
        reporting_period: {
          semester: '1', // TODO: Get from form
          academic_year: '2024-2025', // TODO: Get from form
          start_date: new Date().toISOString().split('T')[0],
          end_date: new Date().toISOString().split('T')[0],
        },
        narrative_content: formData.narrative_content || {
          introduction: '',
          academic_progress: '',
          behavioral_observations: '',
          strengths: [],
          areas_for_improvement: [],
          conclusion: '',
        },
      };

      const report = await createNarrativeReport(reportData);
      navigate(`/narrative-reports/${report.id}`);
    } catch (error) {
      setError('Failed to create narrative report');
      console.error('Failed to create narrative report:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Container maxWidth="xl" sx={{ py: 4 }}>
      <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
        <Link underline="hover" onClick={() => navigate('/narrative-reports')}>
          Narrative Reports
        </Link>
        <Typography color="text.primary">Generate Report</Typography>
      </Breadcrumbs>

      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" component="h1">
          Generate Narrative Report
        </Typography>
        <Button
          variant="outlined"
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/narrative-reports')}
        >
          Back to Reports
        </Button>
      </Box>

      <Paper sx={{ p: 3 }}>
        {error && (
          <Alert severity="error" sx={{ mb: 3 }}>
            {error}
          </Alert>
        )}

        <form onSubmit={handleSubmit}>
          <Grid container spacing={3}>
            <Grid size={{ xs: 12, md: 6 }}>
              <TextField
                label="Student ID"
                name="student_id"
                value={formData.student_id}
                onChange={handleInputChange}
                fullWidth
                required
                helperText="Enter the student ID to generate a report for"
              />
            </Grid>
            <Grid size={{ xs: 12, md: 6 }}>
              <TextField
                label="Period ID"
                name="period_id"
                value={formData.period_id}
                onChange={handleInputChange}
                fullWidth
                helperText="Enter the period ID for this report"
              />
            </Grid>
            <Grid size={{ xs: 12 }}>
              <Button
                variant="outlined"
                onClick={handleLoadAchievementData}
                disabled={isLoading || !formData.student_id}
              >
                {isLoading ? 'Loading...' : 'Load Achievement Data'}
              </Button>
            </Grid>
            <Grid size={{ xs: 12 }}>
              <Divider sx={{ my: 2 }} />
              <Typography variant="h6" gutterBottom>
                Narrative Content
              </Typography>
              <TextField
                name="narrative_content"
                value={formData.narrative_content}
                onChange={handleInputChange}
                fullWidth
                multiline
                rows={10}
                placeholder="Enter the narrative content for this report. This should include observations about the student's progress, achievements, and areas for improvement."
                helperText="Provide a detailed narrative about the student's learning progress"
              />
            </Grid>
            {achievementData && (
              <Grid size={{ xs: 12 }}>
                <Paper sx={{ p: 2, mt: 2, bgcolor: 'background.default' }}>
                  <Typography variant="subtitle1" gutterBottom>
                    Achievement Data Preview
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Total Achievements: {achievementData.total_achievements || 0}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    Average Mastery: {achievementData.average_mastery || 0}%
                  </Typography>
                </Paper>
              </Grid>
            )}
            <Grid size={{ xs: 12 }}>
              <Divider sx={{ my: 2 }} />
              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                <Button
                  variant="outlined"
                  onClick={() => navigate('/narrative-reports')}
                  disabled={isLoading}
                >
                  Cancel
                </Button>
                <Button
                  type="submit"
                  variant="contained"
                  disabled={isLoading || !formData.student_id || !formData.narrative_content}
                >
                  {isLoading ? 'Creating...' : 'Create Report'}
                </Button>
              </Box>
            </Grid>
          </Grid>
        </form>
      </Paper>
    </Container>
  );
};

export default NarrativeReportGeneratePage;
