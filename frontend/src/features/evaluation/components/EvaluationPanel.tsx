/**
 * Evaluation Panel Component
 * Panel for viewing and managing evaluation details
 */

import { Box, Typography, Paper, Divider, Button, Stack } from '@mui/material';
import { Evaluation } from '@/api/evaluation';

interface EvaluationPanelProps {
  evaluation: Evaluation;
  onEdit?: () => void;
  onDelete?: () => void;
  canEdit?: boolean;
  canDelete?: boolean;
}

export const EvaluationPanel = ({
  evaluation,
  onEdit,
  onDelete,
  canEdit = true,
  canDelete = true,
}: EvaluationPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Evaluation Details
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Typography variant="body1">{evaluation.status}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Evidence ID
          </Typography>
          <Typography variant="body1">{evaluation.evidence_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Evaluator ID
          </Typography>
          <Typography variant="body1">{evaluation.evaluator_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Score
          </Typography>
          <Typography variant="body1">{evaluation.score}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Feedback
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {evaluation.feedback || 'No feedback provided'}
          </Typography>
        </Box>

        <Divider />

        <Stack direction="row" spacing={2} justifyContent="flex-end">
          {onEdit && canEdit && (
            <Button variant="outlined" onClick={onEdit}>
              Edit
            </Button>
          )}
          {onDelete && canDelete && (
            <Button variant="contained" color="error" onClick={onDelete}>
              Delete
            </Button>
          )}
        </Stack>
      </Stack>
    </Paper>
  );
};
