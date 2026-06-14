import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Box, Button, Container, Paper, Typography, Breadcrumbs, Link, CircularProgress, Alert } from '@mui/material';
import { CPDetail } from '@/features/cp';
import { useCPById, useSubjects, usePhases, useElementsByPhase, useSubelementsByElement } from '@/services/queries/CPQueryService';
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

const CPDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const { data: cp, isLoading, error } = useCPById(id || '');
  const { data: subjects = [] } = useSubjects();
  const { data: phases = [] } = usePhases();
  const { data: elements = [] } = useElementsByPhase(cp?.phase_id || '');
  const { data: subelements = [] } = useSubelementsByElement(cp?.element_id || '');

  const handleCreateTPFromCP = (cp: any) => {
    navigate('/tp/create', { state: { selectedCP: cp } });
  };

  const handleBack = () => {
    navigate('/cp');
  };

  if (isLoading) {
    return (
      <Container maxWidth="xl">
        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '50vh' }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error">
          Failed to load CP details. Please try again.
        </Alert>
      </Container>
    );
  }

  if (!cp) {
    return (
      <Container maxWidth="xl">
        <Alert severity="info">
          CP not found.
        </Alert>
      </Container>
    );
  }

  const subject = subjects.find((s: any) => s.id === cp.subject_id);
  const phase = phases.find((p: any) => p.id === cp.phase_id);
  const element = elements.find((e: any) => e.id === cp.element_id);
  const subelement = subelements.find((s: any) => s.id === cp.subelement_id);

  return (
    <Container maxWidth="xl">
      <Box sx={{ py: 3 }}>
        <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
          <Link underline="hover" color="inherit" href="/curriculum">
            Curriculum
          </Link>
          <Link underline="hover" color="inherit" onClick={handleBack}>
            CP
          </Link>
          <Typography color="text.primary">{cp.cp_code}</Typography>
        </Breadcrumbs>

        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
          <Typography variant="h4" component="h1">
            CP Detail
          </Typography>
          <Box sx={{ display: 'flex', gap: 2 }}>
            <Button
              variant="outlined"
              startIcon={<ArrowBackIcon />}
              onClick={handleBack}
            >
              Back to CP List
            </Button>
            <Button
              variant="contained"
              startIcon={<NavigateNextIcon />}
              onClick={() => handleCreateTPFromCP(cp)}
            >
              Create TP from this CP
            </Button>
          </Box>
        </Box>

        <Paper sx={{ p: 3 }}>
          <CPDetail
            cp={cp}
            subject={subject}
            phase={phase}
            element={element}
            subelement={subelement}
            onSelectForTP={handleCreateTPFromCP}
          />
        </Paper>
      </Box>
    </Container>
  );
};

export default CPDetailPage;
