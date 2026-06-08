/**
 * Evidence Upload Panel Component
 * Panel for uploading new evidence
 */

import { Box, Typography, Button, Stack, TextField } from '@mui/material';
import { CloudUpload as CloudUploadIcon } from '@mui/icons-material';

interface EvidenceUploadPanelProps {
  onUpload: (file: File, data: any) => void;
  onCancel?: () => void;
  loading?: boolean;
}

export const EvidenceUploadPanel = ({ onUpload, onCancel, loading = false }: EvidenceUploadPanelProps) => {
  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      onUpload(file, { title: file.name });
    }
  };

  return (
    <Box sx={{ p: 3, bgcolor: 'background.paper', borderRadius: 1 }}>
      <Typography variant="h6" gutterBottom>
        Upload Evidence
      </Typography>
      <Stack spacing={3}>
        <Box
          sx={{
            border: '2px dashed',
            borderColor: 'divider',
            borderRadius: 1,
            p: 4,
            textAlign: 'center',
            cursor: 'pointer',
            '&:hover': {
              borderColor: 'primary.main',
            },
          }}
          onClick={() => document.getElementById('evidence-file-input')?.click()}
        >
          <CloudUploadIcon sx={{ fontSize: 48, color: 'text.secondary', mb: 2 }} />
          <Typography variant="body1" color="text.secondary">
            Click to upload or drag and drop
          </Typography>
          <Typography variant="caption" color="text.secondary">
            Supported formats: PDF, DOC, DOCX, Images
          </Typography>
          <input
            id="evidence-file-input"
            type="file"
            accept=".pdf,.doc,.docx,.jpg,.jpeg,.png"
            style={{ display: 'none' }}
            onChange={handleFileSelect}
          />
        </Box>

        <TextField
          fullWidth
          label="Evidence Title"
          placeholder="Enter a title for the evidence"
          size="small"
        />

        <TextField
          fullWidth
          multiline
          rows={3}
          label="Description"
          placeholder="Describe the evidence"
          size="small"
        />

        <Stack direction="row" spacing={2} justifyContent="flex-end">
          {onCancel && (
            <Button variant="outlined" onClick={onCancel} disabled={loading}>
              Cancel
            </Button>
          )}
          <Button variant="contained" disabled={loading}>
            {loading ? 'Uploading...' : 'Upload'}
          </Button>
        </Stack>
      </Stack>
    </Box>
  );
};
