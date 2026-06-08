import React from 'react';
import { Box, Typography, Timeline, TimelineItem, TimelineSeparator, TimelineConnector, TimelineContent, TimelineDot, Chip } from '@mui/material';
import { History as HistoryIcon } from '@mui/icons-material';

interface EvaluationRevision {
  id: string;
  revision_no: number;
  total_score: number;
  performance_level: string;
  teacher_feedback: string;
  created_at: string;
  is_current_version: boolean;
}

interface EvaluationHistoryProps {
  evidenceId: string;
}

export const EvaluationHistory: React.FC<EvaluationHistoryProps> = ({ evidenceId }) => {
  const [history, setHistory] = React.useState<EvaluationRevision[]>([]);
  const [loading, setLoading] = React.useState(true);

  React.useEffect(() => {
    // Fetch evaluation history
    fetch(`/api/v1/assessment/evaluations/history/${evidenceId}`)
      .then(res => res.json())
      .then(data => setHistory(data.data))
      .finally(() => setLoading(false));
  }, [evidenceId]);

  if (loading) return <Box>Loading...</Box>;

  return (
    <Box>
      <Typography variant="h6" gutterBottom>Evaluation History</Typography>
      <Timeline>
        {history.map((revision) => (
          <TimelineItem key={revision.id}>
            <TimelineSeparator>
              <TimelineDot color={revision.is_current_version ? "primary" : "grey"} />
              <TimelineConnector />
            </TimelineSeparator>
            <TimelineContent>
              <Typography variant="body1">Revision {revision.revision_no}</Typography>
              <Typography variant="body2" color="textSecondary">
                Score: {revision.total_score} - {revision.performance_level}
              </Typography>
              <Typography variant="body2">{revision.teacher_feedback}</Typography>
              {revision.is_current_version && (
                <Chip label="Current" color="primary" size="small" />
              )}
            </TimelineContent>
          </TimelineItem>
        ))}
      </Timeline>
    </Box>
  );
};
