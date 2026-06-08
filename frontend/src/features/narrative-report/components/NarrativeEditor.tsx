/**
 * Narrative Editor Component
 * Rich text editor for narrative report content
 */

import { Box, TextField, Typography } from '@mui/material';

interface NarrativeEditorProps {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
}

export const NarrativeEditor = ({ value, onChange, placeholder = 'Start writing your narrative...' }: NarrativeEditorProps) => {
  return (
    <Box>
      <Typography variant="subtitle2" gutterBottom>
        Narrative Content
      </Typography>
      <TextField
        fullWidth
        multiline
        rows={12}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
        sx={{
          '& .MuiInputBase-root': {
            fontFamily: 'monospace',
            fontSize: '14px',
          },
        }}
      />
    </Box>
  );
};
