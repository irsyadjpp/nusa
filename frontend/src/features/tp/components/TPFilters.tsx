/**
 * TP Filters Component
 * Provides filtering options for TP list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface TPFiltersProps {
  filters: {
    subject_id?: string;
    phase_id?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  subjects?: { id: string; name: string }[];
  phases?: { id: string; name: string }[];
}

export const TPFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  subjects = [],
  phases = [],
}: TPFiltersProps) => {
  return (
    <Box sx={{ mb: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Stack direction="row" spacing={2} alignItems="center">
        <TextField
          select
          label="Subject"
          value={filters.subject_id || ''}
          onChange={(e) => onFilterChange({ ...filters, subject_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Subjects</MenuItem>
          {subjects.map((subject) => (
            <MenuItem key={subject.id} value={subject.id}>
              {subject.name}
            </MenuItem>
          ))}
        </TextField>

        <TextField
          select
          label="Phase"
          value={filters.phase_id || ''}
          onChange={(e) => onFilterChange({ ...filters, phase_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Phases</MenuItem>
          {phases.map((phase) => (
            <MenuItem key={phase.id} value={phase.id}>
              {phase.name}
            </MenuItem>
          ))}
        </TextField>

        <TextField
          select
          label="Status"
          value={filters.status || ''}
          onChange={(e) => onFilterChange({ ...filters, status: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 120 }}
        >
          <MenuItem value="">All Status</MenuItem>
          <MenuItem value="draft">Draft</MenuItem>
          <MenuItem value="pending">Pending</MenuItem>
          <MenuItem value="approved">Approved</MenuItem>
        </TextField>

        <Button onClick={onClearFilters} size="small">
          Clear
        </Button>
      </Stack>
    </Box>
  );
};
