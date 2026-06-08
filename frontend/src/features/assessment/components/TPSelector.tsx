/**
 * TP Selector Component
 * Allows selection of a Teaching Plan for assessment creation
 */

import { Box, TextField, MenuItem, Typography } from '@mui/material';

interface TPSelectorProps {
  tps: { id: string; title: string; version_no: number }[];
  selectedTPId?: string;
  onSelect: (tpId: string) => void;
  label?: string;
}

export const TPSelector = ({ tps, selectedTPId, onSelect, label = 'Select TP' }: TPSelectorProps) => {
  return (
    <Box>
      <Typography variant="subtitle2" gutterBottom>
        {label}
      </Typography>
      <TextField
        fullWidth
        select
        value={selectedTPId || ''}
        onChange={(e) => onSelect(e.target.value)}
        size="small"
      >
        <MenuItem value="">Select a TP</MenuItem>
        {tps.map((tp) => (
          <MenuItem key={tp.id} value={tp.id}>
            {tp.title} (v{tp.version_no})
          </MenuItem>
        ))}
      </TextField>
    </Box>
  );
};
