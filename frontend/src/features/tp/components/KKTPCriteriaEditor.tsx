/**
 * KKTP Criteria Editor Component
 * Editor for editing KKTP (Kriteria Ketuntasan Tujuan Pembelajaran) criteria
 */

import { Box, Typography, TextField, Button, Stack, IconButton } from '@mui/material';
import { Add as AddIcon, Delete as DeleteIcon } from '@mui/icons-material';

interface Criteria {
  id: string;
  description: string;
  weight: number;
}

interface KKTPCriteriaEditorProps {
  criteria: Criteria[];
  onChange: (criteria: Criteria[]) => void;
}

export const KKTPCriteriaEditor = ({ criteria, onChange }: KKTPCriteriaEditorProps) => {
  const addCriterion = () => {
    const newCriteria: Criteria = {
      id: Date.now().toString(),
      description: '',
      weight: 0,
    };
    onChange([...criteria, newCriteria]);
  };

  const updateCriteria = (id: string, field: keyof Criteria, value: string | number) => {
    onChange(
      criteria.map((c) => (c.id === id ? { ...c, [field]: value } : c))
    );
  };

  const removeCriteria = (id: string) => {
    onChange(criteria.filter((c) => c.id !== id));
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        KKTP Criteria
      </Typography>
      <Stack spacing={2}>
        {criteria.map((criterion, index) => (
          <Box
            key={criterion.id}
            sx={{
              p: 2,
              border: '1px solid',
              borderColor: 'divider',
              borderRadius: 1,
            }}
          >
            <Stack direction="row" spacing={2} alignItems="flex-start">
              <Box sx={{ flexGrow: 1 }}>
                <Typography variant="subtitle2" gutterBottom>
                  Criterion {index + 1}
                </Typography>
                <TextField
                  fullWidth
                  multiline
                  rows={2}
                  label="Description"
                  value={criterion.description}
                  onChange={(e) => updateCriteria(criterion.id, 'description', e.target.value)}
                  size="small"
                  sx={{ mb: 1 }}
                />
                <TextField
                  type="number"
                  label="Weight"
                  value={criterion.weight}
                  onChange={(e) => updateCriteria(criterion.id, 'weight', parseFloat(e.target.value) || 0)}
                  size="small"
                  InputProps={{ inputProps: { min: 0, max: 100 } }}
                />
              </Box>
              <IconButton onClick={() => removeCriteria(criterion.id)} color="error">
                <DeleteIcon />
              </IconButton>
            </Stack>
          </Box>
        ))}
        <Button
          variant="outlined"
          startIcon={<AddIcon />}
          onClick={addCriterion}
          sx={{ alignSelf: 'flex-start' }}
        >
          Add Criterion
        </Button>
      </Stack>
    </Box>
  );
};
