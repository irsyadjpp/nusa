/**
 * Evaluation Filters Component
 * Provides filtering options for Evaluation list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface EvaluationFiltersProps {
  filters: {
    evidence_id?: string;
    evaluator_id?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  evidences?: { id: string; title: string }[];
  evaluators?: { id: string; name: string }[];
}

export const EvaluationFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  evidences = [],
  evaluators = [],
}: EvaluationFiltersProps) => {
  return (
    <Box sx={{ mb: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Stack direction="row" spacing={2} alignItems="center">
        <TextField
          select
          label="Evidence"
          value={filters.evidence_id || ''}
          onChange={(e) => onFilterChange({ ...filters, evidence_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Evidences</MenuItem>
          {evidences.map((evidence) => (
            <MenuItem key={evidence.id} value={evidence.id}>
              {evidence.title}
            </MenuItem>
          ))}
        </TextField>

        <TextField
          select
          label="Evaluator"
          value={filters.evaluator_id || ''}
          onChange={(e) => onFilterChange({ ...filters, evaluator_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Evaluators</MenuItem>
          {evaluators.map((evaluator) => (
            <MenuItem key={evaluator.id} value={evaluator.id}>
              {evaluator.name}
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
