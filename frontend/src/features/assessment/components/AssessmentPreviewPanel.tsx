/**
 * Assessment Preview Panel Component
 * Displays a preview of the assessment
 */

import { Box, Typography, Paper, Divider, Chip } from '@mui/material';
import { Assessment } from '@/shared/types/domain';

interface AssessmentPreviewPanelProps {
  assessment: Assessment;
}

export const AssessmentPreviewPanel = ({ assessment }: AssessmentPreviewPanelProps) => {
  return (
    <Paper elevation={2} sx={{ p: 3 }}>
      <Typography variant="h5" gutterBottom>
        Assessment Preview
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Status
          </Typography>
          <Chip
            label={assessment.status}
            size="small"
            color={assessment.status === 'APPROVED' ? 'success' : assessment.status === 'DRAFT' ? 'default' : 'warning'}
          />
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Version
          </Typography>
          <Typography variant="body1">{assessment.version_no}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            Assessment Type
          </Typography>
          <Typography variant="body1">{assessment.assessment_type}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            TP ID
          </Typography>
          <Typography variant="body1">{assessment.tp_id}</Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary">
            TP Version
          </Typography>
          <Typography variant="body1">{assessment.tp_version_no}</Typography>
        </Box>

        <Divider />

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Assessment Items
          </Typography>
          <Typography variant="body2" whiteSpace="pre-wrap">
            {typeof assessment.assessment_items === 'string'
              ? assessment.assessment_items
              : JSON.stringify(assessment.assessment_items, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Answer Key
          </Typography>
          <Typography variant="body2" whiteSpace="pre-wrap">
            {typeof assessment.answer_key === 'string'
              ? assessment.answer_key
              : JSON.stringify(assessment.answer_key, null, 2)}
          </Typography>
        </Box>

        <Box>
          <Typography variant="subtitle2" color="text.secondary" gutterBottom>
            Scoring Guidelines
          </Typography>
          <Typography variant="body2" whiteSpace="pre-wrap">
            {typeof assessment.scoring_guidelines === 'string'
              ? assessment.scoring_guidelines
              : JSON.stringify(assessment.scoring_guidelines, null, 2)}
          </Typography>
        </Box>

        {assessment.ai_confidence_score && (
          <Box>
            <Typography variant="subtitle2" color="text.secondary">
              AI Confidence Score
            </Typography>
            <Typography variant="body1">{assessment.ai_confidence_score}%</Typography>
          </Box>
        )}
      </Box>
    </Paper>
  );
};
