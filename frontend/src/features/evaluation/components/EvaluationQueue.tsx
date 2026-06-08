/**
 * Evaluation Queue Component
 * Displays a queue of evaluations pending review
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography, Badge } from '@mui/material';
import { Evaluation } from '@/api/evaluation';

interface EvaluationQueueProps {
  evaluations: Evaluation[];
  selectedId?: string;
  onSelect?: (evaluation: Evaluation) => void;
  loading?: boolean;
}

export const EvaluationQueue = ({ evaluations, selectedId, onSelect, loading }: EvaluationQueueProps) => {
  const pendingCount = evaluations.filter(e => e.status === 'pending').length;

  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading Evaluations...</Typography>
      </Box>
    );
  }

  if (evaluations.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No Evaluations found</Typography>
      </Box>
    );
  }

  return (
    <Box>
      <Box sx={{ mb: 2, display: 'flex', alignItems: 'center', gap: 2 }}>
        <Typography variant="h6">Evaluation Queue</Typography>
        {pendingCount > 0 && (
          <Badge badgeContent={pendingCount} color="error">
            <Typography variant="body2" color="text.secondary">
              Pending
            </Typography>
          </Badge>
        )}
      </Box>
      <List>
        {evaluations.map((evaluation) => (
          <ListItem
            key={evaluation.id}
            disablePadding
            selected={selectedId === evaluation.id}
            secondaryAction={
              <Chip
                label={evaluation.status}
                size="small"
                color={evaluation.status === 'approved' ? 'success' : evaluation.status === 'draft' ? 'default' : 'warning'}
              />
            }
          >
            <ListItemButton onClick={() => onSelect?.(evaluation)} selected={selectedId === evaluation.id}>
              <ListItemText
                primary={`Evidence: ${evaluation.evidence_id}`}
                secondary={`Score: ${evaluation.score} | Evaluator: ${evaluation.evaluator_id}`}
              />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Box>
  );
};
