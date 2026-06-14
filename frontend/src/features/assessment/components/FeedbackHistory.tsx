import React from 'react';
import { Box, Typography, List, ListItem, ListItemText, Divider } from '@mui/material';

interface FeedbackEntry {
  id: string;
  teacher_feedback: string;
  changed_by: string;
  changed_at: string;
}

interface FeedbackHistoryProps {
  evaluationId: string;
}

export const FeedbackHistory: React.FC<FeedbackHistoryProps> = ({ evaluationId }) => {
  const [history, setHistory] = React.useState<FeedbackEntry[]>([]);
  const [loading, setLoading] = React.useState(true);

  React.useEffect(() => {
    fetch(`/api/v1/assessment/evaluations/${evaluationId}/feedback-history`)
      .then(res => res.json())
      .then(data => setHistory(data.data))
      .finally(() => setLoading(false));
  }, [evaluationId]);

  if (loading) return <Box>Loading...</Box>;

  return (
    <Box>
      <Typography variant="h6" gutterBottom>Feedback History</Typography>
      <List>
        {history.map((entry, index) => (
          <React.Fragment key={entry.id}>
            <ListItem>
              <ListItemText
                primary={entry.teacher_feedback}
                secondary={`By: ${entry.changed_by} - ${new Date(entry.changed_at).toLocaleString()}`}
              />
            </ListItem>
            {index < history.length - 1 && <Divider />}
          </React.Fragment>
        ))}
      </List>
    </Box>
  );
};
