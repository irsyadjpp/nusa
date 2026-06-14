/**
 * Evaluation Preview Component
 * Displays a preview of the evaluation
 */

import { Box, Typography, Paper, Divider, Chip, LinearProgress } from '@mui/material';
import { Evaluation } from '@/shared/types/domain';

interface EvaluationPreviewProps {
  evaluation: Evaluation;
}

export const EvaluationPreview = ({ evaluation }: EvaluationPreviewProps) => {
  const scorePercentage = evaluation.performance_scores.max_score > 0 ? (evaluation.performance_scores.total_score / evaluation.performance_scores.max_score) * 100 : 0;

  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Evaluation Preview
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Performance Level
          </Typography>
          <Chip
            label={evaluation.performance_level}
            size="small"
            color={evaluation.performance_level === 'PROFICIENT' ? 'success' : evaluation.performance_level === 'DEVELOPING' ? 'warning' : 'default'}
          />
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Score
          </Typography>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
            <Typography variant="body1">{evaluation.performance_scores.total_score}/{evaluation.performance_scores.max_score}</Typography>
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
          <Typography variant="body1">{evaluation.teacher_id}</Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Feedback
          </Typography>
          <Typography variant="body1" whiteSpace="pre-wrap">
            {evaluation.teacher_feedback || 'No feedback provided'}
          </Typography>
        </Box>

        {evaluation.performance_scores && (
          <Box>
            <Typography variant="subtitle2" color="text.secondary" gutterBottom>
              Performance Scores
            </Typography>
            <Typography variant="body2" whiteSpace="pre-wrap">
              {typeof evaluation.performance_scores === 'string'
                ? evaluation.performance_scores
                : JSON.stringify(evaluation.performance_scores, null, 2)}
            </Typography>
          </Box>
        )}
      </Box>
    </Paper>
  );
};
