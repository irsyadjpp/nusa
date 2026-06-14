/**
 * Evaluation Panel Component
 * Panel for viewing and managing evaluation details
 */

import { Box, Typography, Paper, Divider, Button } from '@mui/material';
import { Evaluation } from '@/shared/types/domain';

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
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
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
          <Typography variant="body1">{evaluation.teacher_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Total Score
          </Typography>
          <Typography variant="body1">{evaluation.performance_scores.total_score} / {evaluation.performance_scores.max_score}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Performance Level
          </Typography>
          <Typography variant="body1">{evaluation.performance_level}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Feedback
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {evaluation.teacher_feedback || 'No feedback provided'}
          </Typography>
        </Box>

        <Divider />

        <Box sx={{ display: 'flex', flexDirection: 'row', gap: 2, justifyContent: 'flex-end' }}>
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
        </Box>
      </Box>
    </Paper>
  );
};
