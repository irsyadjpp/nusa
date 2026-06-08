/**
 * Evaluation Preview Component
 * Displays a preview of the evaluation
 */

import { Box, Typography, Paper, Divider, Chip, LinearProgress } from '@mui/material';
import { Evaluation } from '@/api/evaluation';

interface EvaluationPreviewProps {
  evaluation: Evaluation;
}

export const EvaluationPreview = ({ evaluation }: EvaluationPreviewProps) => {
  const scorePercentage = evaluation.score || 0;

  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Evaluation Preview
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Stack spacing={2}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Chip
            label={evaluation.status}
            size="small"
            color={evaluation.status === 'approved' ? 'success' : evaluation.status === 'draft' ? 'default' : 'warning'}
          />
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Score
          </Typography>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
            <Typography variant="body1">{evaluation.score}/100</Typography>
            <LinearProgress
              variant="determinate"
              value={scorePercentage}
              sx={{ flexGrow: 1 }}
              color={scorePercentage >= 70 ? 'success' : scorePercentage >= 50 ? 'warning' : 'error'}
            />
          </Box>
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

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Feedback
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {evaluation.feedback || 'No feedback provided'}
          </Typography>
        </Box>

        {evaluation.rubric_scores && (
          <Box>
            <Typography variant="subtitle2" color="text.secondary" gutterBottom>
              Rubric Scores
            </Typography>
            <Typography variant="body2" whiteSpace="pre-wrap">
              {typeof evaluation.rubric_scores === 'string'
                ? evaluation.rubric_scores
                : JSON.stringify(evaluation.rubric_scores, null, 2)}
            </Typography>
          </Box>
        )}
      </Stack>
    </Paper>
  );
};
