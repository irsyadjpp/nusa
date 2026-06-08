/**
 * Modul Ajar Filters Component
 * Provides filtering options for Modul Ajar list
 */

import { Box, TextField, MenuItem, Stack, Button } from '@mui/material';

interface ModulAjarFiltersProps {
  filters: {
    modul_ajar_set_id?: string;
    atp_id?: string;
    tp_id?: string;
    status?: string;
  };
  onFilterChange: (filters: any) => void;
  onClearFilters: () => void;
  modulAjarSets?: { id: string; name: string }[];
  atps?: { id: string; name: string }[];
  tps?: { id: string; title: string }[];
}

export const ModulAjarFilters = ({
  filters,
  onFilterChange,
  onClearFilters,
  modulAjarSets = [],
  atps = [],
  tps = [],
}: ModulAjarFiltersProps) => {
  return (
    <Box sx={{ mb: 2, p: 2, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Stack direction="row" spacing={2} alignItems="center">
        <TextField
          select
          label="Modul Ajar Set"
          value={filters.modul_ajar_set_id || ''}
          onChange={(e) => onFilterChange({ ...filters, modul_ajar_set_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All Sets</MenuItem>
          {modulAjarSets.map((set) => (
            <MenuItem key={set.id} value={set.id}>
              {set.name}
            </MenuItem>
          ))}
        </TextField>

        <TextField
          select
          label="ATP"
          value={filters.atp_id || ''}
          onChange={(e) => onFilterChange({ ...filters, atp_id: e.target.value || undefined })}
          size="small"
          sx={{ minWidth: 150 }}
        >
          <MenuItem value="">All ATPs</MenuItem>
          {atps.map((atp) => (
            <MenuItem key={atp.id} value={atp.id}>
              {atp.name}
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
