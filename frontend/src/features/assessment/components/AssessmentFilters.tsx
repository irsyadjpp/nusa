/**
 * Assessment Filters Component
 * Provides filtering options for Assessment list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface AssessmentFiltersProps {
  filters: {
    tp_id?: string;
    user_id?: string;
    assessment_type?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  tps?: { id: string; title: string }[];
  users?: { id: string; name: string }[];
}

export const AssessmentFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  tps = [],
  users = [],
}: AssessmentFiltersProps) => {
  return (
    <Box sx={{ mb: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Stack direction="row" spacing={2} alignItems="center">
        <TextField
          select
          label="TP"
          value={filters.tp_id || ''}
          onChange={(e) => onFilterChange({ ...filters, tp_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All TPs</MenuItem>
          {tps.map((tp) => (
            <MenuItem key={tp.id} value={tp.id}>
              {tp.title}
            </MenuItem>
          ))}
        </TextField>

        <TextField
          select
          label="User"
          value={filters.user_id || ''}
          onChange={(e) => onFilterChange({ ...filters, user_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Users</MenuItem>
          {users.map((user) => (
            <MenuItem key={user.id} value={user.id}>
              {user.name}
            </MenuItem>
          ))}
        </TextField>

        <TextField
          select
          label="Assessment Type"
          value={filters.assessment_type || ''}
          onChange={(e) => onFilterChange({ ...filters, assessment_type: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Types</MenuItem>
          <MenuItem value="formative">Formative</MenuItem>
          <MenuItem value="summative">Summative</MenuItem>
          <MenuItem value="diagnostic">Diagnostic</MenuItem>
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
