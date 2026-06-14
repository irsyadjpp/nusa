/**
 * CP (Curriculum Plan) Components
 * Contains all CP-related UI components
 */

import React, { useState } from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  CardActions,
  Button,
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Chip,
  Grid,
  CircularProgress,
  Alert,
  Divider,
} from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import FilterListIcon from '@mui/icons-material/FilterList';
import { CP, CurriculumSubject, CurriculumPhase, CurriculumElement, CurriculumSubelement } from '@/shared/types/domain';
import { useCPs, useSubjects, usePhases, useElementsByPhase, useSubelementsByElement } from '@/services/queries/CPQueryService';

// Type aliases for consistency
type Subject = CurriculumSubject;
type Phase = CurriculumPhase;
type Element = CurriculumElement;
type Subelement = CurriculumSubelement;

// CP Filter Component
interface CPFiltersProps {
  filters: {
    subject_id?: string;
    phase_id?: string;
    search?: string;
  };
  subjects: Subject[];
  phases: Phase[];
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
}

export const CPFilters: React.FC<CPFiltersProps> = ({
  filters,
  subjects,
  phases,
  onFilterChange,
  onClearFilters,
}) => {
  return (
    <Box>
      <Typography variant="h6" gutterBottom sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
        <FilterListIcon fontSize="small" />
        Filters
      </Typography>
      
      <Grid container spacing={2}>
        <Grid size={{ xs: 12, md: 6 }}>
          <TextField
            fullWidth
            size="small"
            label="Search CP"
            placeholder="Search by description or code..."
            value={filters.search || ''}
            onChange={(e) => onFilterChange({ ...filters, search: e.target.value })}
            InputProps={{
              startAdornment: <SearchIcon sx={{ mr: 1, color: 'text.secondary' }} />,
            }}
          />
        </Grid>
        
        <Grid size={{ xs: 12, md: 6 }}>
          <FormControl fullWidth size="small">
            <InputLabel>Subject</InputLabel>
            <Select
              value={filters.subject_id || ''}
              label="Subject"
              onChange={(e) => onFilterChange({ ...filters, subject_id: e.target.value || undefined })}
            >
              <MenuItem value="">All Subjects</MenuItem>
              {subjects.map((subject) => (
                <MenuItem key={subject.id} value={subject.id}>
                  {subject.name} ({subject.code})
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        
        <Grid size={{ xs: 12, md: 6 }}>
          <FormControl fullWidth size="small">
            <InputLabel>Phase</InputLabel>
            <Select
              value={filters.phase_id || ''}
              label="Phase"
              onChange={(e) => onFilterChange({ ...filters, phase_id: e.target.value || undefined })}
            >
              <MenuItem value="">All Phases</MenuItem>
              {phases.map((phase) => (
                <MenuItem key={phase.id} value={phase.id}>
                  {phase.name} - {phase.grade_level}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        
        <Grid size={{ xs: 12, md: 6 }}>
          <Button
            fullWidth
            variant="outlined"
            onClick={onClearFilters}
            disabled={!filters.subject_id && !filters.phase_id && !filters.search}
          >
            Clear Filters
          </Button>
        </Grid>
      </Grid>
    </Box>
  );
};

// CP Card Component
interface CPCardProps {
  cp: CP;
  subject?: Subject;
  phase?: Phase;
  element?: Element;
  subelement?: Subelement;
  onSelect?: (cp: CP) => void;
  selected?: boolean;
}

export const CPCard: React.FC<CPCardProps> = ({
  cp,
  subject,
  phase,
  element,
  subelement,
  onSelect,
  selected = false,
}) => {
  return (
    <Card
      sx={{
        cursor: onSelect ? 'pointer' : 'default',
        border: selected ? 2 : 1,
        borderColor: selected ? 'primary.main' : 'divider',
        '&:hover': onSelect ? { borderColor: 'primary.main' } : {},
      }}
      onClick={() => onSelect && onSelect(cp)}
    >
      <CardContent>
        <Box sx={{ mb: 2 }}>
          <Chip
            label={cp.cp_code}
            size="small"
            color="primary"
            variant="outlined"
            sx={{ mb: 1 }}
          />
          <Typography variant="body2" color="text.secondary">
            {subject?.name} &gt; {phase?.name} &gt; {element?.name}
            {subelement && ` &gt; ${subelement.name}`}
          </Typography>
        </Box>
        
        <Typography variant="h6" gutterBottom>
          {cp.cp_text}
        </Typography>
        
        <Typography variant="body2" color="text.secondary">
          ID: {cp.id}
        </Typography>
      </CardContent>
      
      {onSelect && (
        <CardActions>
          <Button size="small" variant="contained">
            Select CP
          </Button>
        </CardActions>
      )}
    </Card>
  );
};

// CP List Component
interface CPListProps {
  cps: CP[];
  subjects: Subject[];
  phases: Phase[];
  elements: Element[];
  subelements: Subelement[];
  selectedId?: string;
  onSelect?: (cp: CP) => void;
  loading?: boolean;
  error?: string;
}

export const CPList: React.FC<CPListProps> = ({
  cps,
  subjects,
  phases,
  elements,
  subelements,
  selectedId,
  onSelect,
  loading = false,
  error,
}) => {
  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', p: 4 }}>
        <CircularProgress />
      </Box>
    );
  }

  if (error) {
    return (
      <Alert severity="error" sx={{ mb: 2 }}>
        {error}
      </Alert>
    );
  }

  if (!cps || cps.length === 0) {
    return (
      <Alert severity="info">
        No Curriculum Plans found. Try adjusting your filters.
      </Alert>
    );
  }

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
      {cps.map((cp) => {
        const subject = subjects.find((s) => s.id === cp.subject_id);
        const phase = phases.find((p) => p.id === cp.phase_id);
        const element = elements.find((e) => e.id === cp.element_id);
        const subelement = subelements.find((s) => s.id === cp.subelement_id);

        return (
          <CPCard
            key={cp.id}
            cp={cp}
            subject={subject}
            phase={phase}
            element={element}
            subelement={subelement}
            onSelect={onSelect}
            selected={selectedId === cp.id}
          />
        );
      })}
    </Box>
  );
};

// CP Detail Component
interface CPDetailProps {
  cp: CP;
  subject?: Subject;
  phase?: Phase;
  element?: Element;
  subelement?: Subelement;
  onClose?: () => void;
  onSelectForTP?: (cp: CP) => void;
}

export const CPDetail: React.FC<CPDetailProps> = ({
  cp,
  subject,
  phase,
  element,
  subelement,
  onClose,
  onSelectForTP,
}) => {
  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
        <Typography variant="h5">CP Detail</Typography>
        {onClose && (
          <Button variant="outlined" onClick={onClose}>
            Close
          </Button>
        )}
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ mb: 3 }}>
            <Chip
              label={cp.cp_code}
              size="small"
              color="primary"
              variant="outlined"
              sx={{ mb: 2 }}
            />
            
            <Typography variant="body1" color="text.secondary" gutterBottom>
              <strong>Curriculum Hierarchy:</strong>
            </Typography>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 1, ml: 2 }}>
              <Typography variant="body2">
                Subject: {subject?.name} ({subject?.code})
              </Typography>
              <Typography variant="body2">
                Phase: {phase?.name} - {phase?.grade_level}
              </Typography>
              <Typography variant="body2">
                Element: {element?.name} ({element?.code})
              </Typography>
              {subelement && (
                <Typography variant="body2">
                  Subelement: {subelement.name} ({subelement.code})
                </Typography>
              )}
            </Box>
          </Box>

          <Divider sx={{ my: 2 }} />

          <Typography variant="h6" gutterBottom>
            Description
          </Typography>
          <Typography variant="body1" paragraph>
            {cp.cp_text}
          </Typography>

          <Typography variant="body2" color="text.secondary">
            ID: {cp.id}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Created: {new Date(cp.created_at).toLocaleDateString()}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Updated: {new Date(cp.updated_at).toLocaleDateString()}
          </Typography>
        </CardContent>

        {onSelectForTP && (
          <CardActions>
            <Button
              variant="contained"
              color="primary"
              onClick={() => onSelectForTP(cp)}
            >
              Create TP from this CP
            </Button>
          </CardActions>
        )}
      </Card>
    </Box>
  );
};

// CP Selection Component (for use in TP form)
interface CPSelectorProps {
  selectedCP?: CP;
  onSelect?: (cp: CP) => void;
  disabled?: boolean;
}

export const CPSelector: React.FC<CPSelectorProps> = ({
  selectedCP,
  onSelect,
  disabled = false,
}) => {
  const [open, setOpen] = useState(false);
  const [filters, setFilters] = useState<{ subject_id?: string; phase_id?: string; element_id?: string; search?: string }>({});

  const { data: cps = [], isLoading } = useCPs(filters);
  const { data: subjects = [] } = useSubjects();
  const { data: phases = [] } = usePhases();
  const { data: elements = [] } = useElementsByPhase(filters.phase_id || '');
  const { data: subelements = [] } = useSubelementsByElement(filters.element_id || '');

  return (
    <Box>
      <Typography variant="subtitle2" gutterBottom>
        Select Curriculum Plan (CP)
      </Typography>
      
      <Box sx={{ display: 'flex', gap: 1, alignItems: 'center' }}>
        <TextField
          fullWidth
          size="small"
          placeholder={selectedCP ? `${selectedCP.cp_code}: ${selectedCP.cp_text}` : 'Select a CP...'}
          value={selectedCP ? `${selectedCP.cp_code}: ${selectedCP.cp_text}` : ''}
          disabled
          InputProps={{
            endAdornment: (
              <Button
                size="small"
                variant="outlined"
                onClick={() => setOpen(true)}
                disabled={disabled}
              >
                {selectedCP ? 'Change' : 'Select'}
              </Button>
            ),
          }}
        />
      </Box>

      {open && (
        <Box
          sx={{
            position: 'absolute',
            zIndex: 1300,
            width: '100%',
            maxWidth: 800,
            maxHeight: 600,
            overflow: 'auto',
            bgcolor: 'background.paper',
            boxShadow: 3,
            borderRadius: 1,
            p: 2,
            mt: 1,
          }}
        >
          <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
            <Typography variant="h6">Select Curriculum Plan</Typography>
            <Button size="small" onClick={() => setOpen(false)}>
              Close
            </Button>
          </Box>

          <CPFilters
            filters={filters}
            subjects={subjects}
            phases={phases}
            onFilterChange={setFilters}
            onClearFilters={() => setFilters({})}
          />

          <Divider sx={{ my: 2 }} />

          <CPList
            cps={cps}
            subjects={subjects}
            phases={phases}
            elements={elements}
            subelements={subelements}
            selectedId={selectedCP?.id}
            onSelect={(cp) => {
              onSelect && onSelect(cp);
              setOpen(false);
            }}
            loading={isLoading}
          />
        </Box>
      )}
    </Box>
  );
};
