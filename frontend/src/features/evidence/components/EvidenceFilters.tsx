/**
 * Evidence Filters Component
 * Provides filtering options for Evidence list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface EvidenceFiltersProps {
  filters: {
    student_id?: string;
    assessment_id?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  students?: { id: string; name: string }[];
  assessments?: { id: string; title: string }[];
}

export const EvidenceFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  students = [],
  assessments = [],
}: EvidenceFiltersProps) => {
  return (
    <Box sx={{ mb: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Stack direction="row" spacing={2} alignItems="center">
        <TextField
          select
          label="Student"
          value={filters.student_id || ''}
          onChange={(e) => onFilterChange({ ...filters, student_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Students</MenuItem>
          {students.map((student) => (
            <MenuItem key={student.id} value={student.id}>
              {student.name}
            </MenuItem>
          ))}
        </TextField>

        <TextField
          select
          label="Assessment"
          value={filters.assessment_id || ''}
          onChange={(e) => onFilterChange({ ...filters, assessment_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
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
