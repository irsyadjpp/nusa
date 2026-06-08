/**
 * ATP Filters Component
 * Provides filtering options for ATP list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface ATPFiltersProps {
  filters: {
    atp_set_id?: string;
    tp_id?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  atpSets?: { id: string; name: string }[];
  tps?: { id: string; title: string }[];
}

export const ATPFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  atpSets = [],
  tps = [],
}: ATPFiltersProps) => {
  return (
    <Box sx={{ mb: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Stack direction="row" spacing={2} alignItems="center">
        <TextField
          select
          label="ATP Set"
          value={filters.atp_set_id || ''}
          onChange={(e) => onFilterChange({ ...filters, atp_set_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All ATP Sets</MenuItem>
          {atpSets.map((set) => (
            <MenuItem key={set.id} value={set.id}>
              {set.name}
            </MenuItem>
          ))}
        </TextField>

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
