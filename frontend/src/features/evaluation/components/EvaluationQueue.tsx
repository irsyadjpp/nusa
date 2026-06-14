/**
 * Evaluation Queue Component
 * Displays a queue of evaluations pending review
 */

import React from 'react';
import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography, Badge, Divider } from '@mui/material';
import { Evaluation } from '@/shared/types/domain';

interface EvaluationQueueProps {
  evaluations: Evaluation[];
  selectedId?: string;
  onSelect?: (evaluation: Evaluation) => void;
  loading?: boolean;
}

export const EvaluationQueue = ({ evaluations, selectedId, onSelect, loading }: EvaluationQueueProps) => {
  // Remove pending count since status doesn't exist
  const pendingCount = 0;

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
        {evaluations.map((evaluation, index) => (
          <React.Fragment key={evaluation.id}>
            <ListItem
              disablePadding
              sx={{
                backgroundColor: selectedId === evaluation.id ? 'action.selected' : 'transparent',
              }}
            >
              <ListItemButton onClick={() => onSelect?.(evaluation)}>
                <ListItemText
                  primary={`Evidence: ${evaluation.evidence_id}`}
                  secondary={`Score: ${evaluation.performance_scores.total_score}/${evaluation.performance_scores.max_score} | Level: ${evaluation.performance_level}`}
                />
              </ListItemButton>
              <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
                <Chip
                  label={evaluation.performance_level}
                  size="small"
                  color={evaluation.performance_level === 'PROFICIENT' ? 'success' : evaluation.performance_level === 'DEVELOPING' ? 'warning' : 'default'}
                />
              </Box>
            </ListItem>
            {index < evaluations.length - 1 && <Divider />}
          </React.Fragment>
        ))}
      </List>
    </Box>
  );
};
