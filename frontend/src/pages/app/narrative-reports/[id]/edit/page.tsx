import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import {
  Box,
  Typography,
  Button,
  Paper,
  TextField,
  CircularProgress,
  Alert,
  Container,
  Breadcrumbs,
  Link,
  Grid,
  Divider,
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import {
  getNarrativeReportById,
  updateNarrativeReport,
} from '@/api/narrative-report';
import { NarrativeContent } from '@/shared/types/domain';

const NarrativeReportEditPage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const [formData, setFormData] = useState({
    narrative_content: '',
    status: '',
  });

  const [isLoading, setIsLoading] = useState(true);
  const [isSaving, setIsSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    loadReport();
  }, [id]);

  const loadReport = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const report = await getNarrativeReportById(id!);
      setFormData({
        narrative_content: typeof report.narrative_content === 'string' ? report.narrative_content : JSON.stringify(report.narrative_content),
        status: report.status,
      });
    } catch (error) {
      setError('Failed to load narrative report');
      console.error('Failed to load narrative report:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      setIsSaving(true);
      setError(null);

      let narrativeContent: NarrativeContent;
      try {
        narrativeContent = typeof formData.narrative_content === 'string' 
          ? JSON.parse(formData.narrative_content) as NarrativeContent
          : formData.narrative_content as NarrativeContent;
      } catch {
        narrativeContent = {
          introduction: '',
          academic_progress: '',
          behavioral_observations: '',
          strengths: [],
          areas_for_improvement: [],
          conclusion: '',
        };
      }

      const reportData = {
        narrative_content: narrativeContent,
      };

      await updateNarrativeReport(id!, reportData);
      navigate(`/narrative-reports/${id}`);
    } catch (error) {
      setError('Failed to update narrative report');
      console.error('Failed to update narrative report:', error);
    } finally {
      setIsSaving(false);
    }
  };

  if (isLoading) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl" sx={{ py: 4 }}>
      <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
        <Link underline="hover" onClick={() => navigate('/narrative-reports')}>
          Narrative Reports
        </Link>
        <Link underline="hover" onClick={() => navigate(`/narrative-reports/${id}`)}>
          Report Detail
        </Link>
        <Typography color="text.primary">Edit Report</Typography>
      </Breadcrumbs>

      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" component="h1">
          Edit Narrative Report
        </Typography>
        <Button
          variant="outlined"
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate(`/narrative-reports/${id}`)}
        >
          Back to Report
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
            <Grid size={{ xs: 12 }}>
              <TextField
                label="Status"
                name="status"
                value={formData.status}
                fullWidth
                disabled
                helperText="Current status of the report (cannot be changed here)"
              />
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
            <Grid size={{ xs: 12 }}>
              <Divider sx={{ my: 2 }} />
              <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                <Button
                  variant="outlined"
                  onClick={() => navigate(`/narrative-reports/${id}`)}
                  disabled={isSaving}
                >
                  Cancel
                </Button>
                <Button
                  type="submit"
                  variant="contained"
                  disabled={isSaving || !formData.narrative_content}
                >
                  {isSaving ? 'Saving...' : 'Save Changes'}
                </Button>
              </Box>
            </Grid>
          </Grid>
        </form>
      </Paper>
    </Container>
  );
};

export default NarrativeReportEditPage;
