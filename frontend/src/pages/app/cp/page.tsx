import React, { useState } from 'react';
import { Box, Paper, Button, Typography, Container, Breadcrumbs, Link, Divider } from '@mui/material';
import { CPFilters, CPList, CPDetail } from '@/features/cp';
import { useCPs, useSubjects, usePhases, useElementsByPhase, useSubelementsByElement } from '@/services/queries/CPQueryService';
import { useNavigate } from 'react-router-dom';
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import { Add as AddIcon } from '@mui/icons-material';
import { CP } from '@/shared/types/domain';

const CurriculumPlanPage: React.FC = () => {
  const navigate = useNavigate();
  const [filters, setFilters] = useState<any>({});
  const [selectedCP, setSelectedCP] = useState<CP | null>(null);

  // Query data - using type narrowing and array safety
  const { data: cpsData, isLoading, error } = useCPs(filters);
  const { data: subjectsData } = useSubjects();
  const { data: phasesData } = usePhases();
  const { data: elementsData } = useElementsByPhase((filters as any).phase_id || '');
  const { data: subelementsData } = useSubelementsByElement((filters as any).element_id || '');

  // Ensure data is always an array - context7 best practice for data safety
  const cps = Array.isArray(cpsData) ? cpsData : [];
  const subjects = Array.isArray(subjectsData) ? subjectsData : [];
  const phases = Array.isArray(phasesData) ? phasesData : [];
  const elements = Array.isArray(elementsData) ? elementsData : [];
  const subelements = Array.isArray(subelementsData) ? subelementsData : [];

  const handleFilterChange = (newFilters: any) => {
    setFilters(newFilters);
  };

  const handleClearFilters = () => {
    setFilters({});
  };

  const handleSelectCP = (cp: any) => {
    setSelectedCP(cp);
  };

  const handleCreateTPFromCP = (cp: any) => {
    navigate('/tp/create', { state: { selectedCP: cp } });
  };

  const handleCloseDetail = () => {
    setSelectedCP(null);
  };

  // Get related data for selected CP - context7: use optional chaining and array safety
  const selectedCPSubject = subjects.find((s) => s.id === selectedCP?.subject_id);
  const selectedCPPhase = phases.find((p) => p.id === selectedCP?.phase_id);
  const selectedCPElement = elements.find((e) => e.id === selectedCP?.element_id);
  const selectedCPSubelement = subelements.find((s) => s.id === selectedCP?.subelement_id);

  return (
    <Container maxWidth="xl">
      <Box sx={{ py: 3 }}>
        <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
          <Link underline="hover" color="inherit" href="/curriculum">
            Curriculum
          </Link>
          <Typography color="text.primary">Curriculum Plan (CP)</Typography>
        </Breadcrumbs>

        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
          <Typography variant="h4" component="h1">
            Curriculum Plan (CP)
          </Typography>
          <Box sx={{ display: 'flex', gap: 2 }}>
            <Button
              variant="contained"
              startIcon={<AddIcon />}
              onClick={() => navigate('/cp/create')}
            >
              Buat CP Baru
            </Button>
            <Button
              variant="outlined"
              startIcon={<NavigateNextIcon />}
              onClick={() => navigate('/tp')}
            >
              Go to TP
            </Button>
          </Box>
        </Box>

        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 3 }}>
          {/* Left Panel - CP List */}
          <Box sx={{ flex: selectedCP ? '0 0 50%' : '0 0 100%', minWidth: selectedCP ? '48%' : '100%' }}>
            <Paper sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
              <Box sx={{ p: 2 }}>
                <CPFilters
                  filters={filters}
                  subjects={subjects}
                  phases={phases}
                  onFilterChange={handleFilterChange}
                  onClearFilters={handleClearFilters}
                />
              </Box>

              <Divider />

              <Box sx={{ flex: 1, overflow: 'auto', p: 2 }}>
                <CPList
                  cps={cps}
                  subjects={subjects}
                  phases={phases}
                  elements={elements}
                  subelements={subelements}
                  selectedId={selectedCP?.id}
                  onSelect={handleSelectCP}
                  loading={isLoading}
                  error={typeof error === 'string' ? error : error ? String(error) : undefined}
                />
              </Box>
            </Paper>
          </Box>

          {/* Right Panel - CP Detail */}
          {selectedCP && (
            <Box sx={{ flex: '0 0 50%', minWidth: '48%' }}>
              <Paper sx={{ height: '100%', p: 2 }}>
                <CPDetail
                  cp={selectedCP}
                  subject={selectedCPSubject}
                  phase={selectedCPPhase}
                  element={selectedCPElement}
                  subelement={selectedCPSubelement}
                  onClose={handleCloseDetail}
                  onSelectForTP={handleCreateTPFromCP}
                />
              </Paper>
            </Box>
          )}
        </Box>
      </Box>
    </Container>
  );
};

export default CurriculumPlanPage;
