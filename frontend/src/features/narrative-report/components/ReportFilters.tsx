/**
 * Report Filters Component
 * Provides filtering options for Report list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface ReportFiltersProps {
  filters: {
    student_id?: string;
    period?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  students?: { id: string; name: string }[];
}

export const ReportFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  students = [],
}: ReportFiltersProps) => {
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
          label="Period"
          value={filters.period || ''}
          onChange={(e) => onFilterChange({ ...filters, period: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        />

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
          <MenuItem value="published">Published</MenuItem>
        </TextField>

        <Button onClick={onClearFilters} size="small">
          Clear
        </Button>
      </Stack>
    </Box>
  );
};
