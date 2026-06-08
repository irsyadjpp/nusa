/**
 * Rubric Filters Component
 * Provides filtering options for Rubric list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface RubricFiltersProps {
  filters: {
    assessment_id?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  assessments?: { id: string; title: string }[];
}

export const RubricFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  assessments = [],
}: RubricFiltersProps) => {
  return (
    <Box sx={{ mb: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Stack direction="row" spacing={2} alignItems="center">
        <TextField
          select
          label="Assessment"
          value={filters.assessment_id || ''}
          onChange={(e) => onFilterChange({ ...filters, assessment_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 200 }}
        >
          <MenuItem value="">All Assessments</MenuItem>
          {assessments.map((assessment) => (
            <MenuItem key={assessment.id} value={assessment.id}>
              {assessment.title}
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
