/**
 * Modul Ajar Preview Panel Component
 * Displays a preview of the Modul Ajar
 */

import { Box, Typography, Paper, Divider, Chip } from '@mui/material';
import { ModulAjar } from '@/api/modul-ajar';

interface ModulAjarPreviewPanelProps {
  modulAjar: ModulAjar;
}

export const ModulAjarPreviewPanel = ({ modulAjar }: ModulAjarPreviewPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        {modulAjar.title}
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Chip
            label={modulAjar.status}
            size="small"
            color={modulAjar.status === 'approved' ? 'success' : modulAjar.status === 'draft' ? 'default' : 'warning'}
          />
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Sequence Number
          </Typography>
          <Typography variant="body1">{modulAjar.sequence_number}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            ATP ID
          </Typography>
          <Typography variant="body1">{modulAjar.atp_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            TP ID
          </Typography>
          <Typography variant="body1">{modulAjar.tp_id}</Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Learning Activities
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {typeof modulAjar.learning_activities === 'string'
              ? modulAjar.learning_activities
              : JSON.stringify(modulAjar.learning_activities, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Teaching Methods
          </Typography>
          <Typography variant="body1">
            {Array.isArray(modulAjar.teaching_methods)
              ? modulAjar.teaching_methods.join(', ')
              : JSON.stringify(modulAjar.teaching_methods, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Learning Media
          </Typography>
          <Typography variant="body1">
            {Array.isArray(modulAjar.learning_media)
              ? modulAjar.learning_media.join(', ')
              : JSON.stringify(modulAjar.learning_media, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Learning Resources
          </Typography>
          <Typography variant="body1">
            {Array.isArray(modulAjar.learning_resources)
              ? modulAjar.learning_resources.join(', ')
              : JSON.stringify(modulAjar.learning_resources, null, 2)}
          </Typography>
        </Box>
      </Stack>
    </Paper>
  );
};
