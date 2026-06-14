/**
 * Evidence List Component
 * Displays a list of evidences with filtering and pagination
 */

import React from 'react';
import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography, Divider } from '@mui/material';
import { Evidence } from '@/shared/types/domain';

interface EvidenceListProps {
  evidences: Evidence[];
  selectedId?: string;
  onSelect?: (evidence: Evidence) => void;
  loading?: boolean;
}

export const EvidenceList = ({ evidences, selectedId, onSelect, loading }: EvidenceListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading Evidences...</Typography>
      </Box>
    );
  }

  if (evidences.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No Evidences found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {evidences.map((evidence, index) => (
        <React.Fragment key={evidence.id}>
          <ListItem
            disablePadding
            sx={{
              backgroundColor: selectedId === evidence.id ? 'action.selected' : 'transparent',
            }}
          >
            <ListItemButton onClick={() => onSelect?.(evidence)}>
              <ListItemText
                primary={evidence.title || `Evidence ${evidence.id}`}
                secondary={`Student: ${evidence.student_id} | Assessment: ${evidence.assessment_id}`}
              />
            </ListItemButton>
            <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
              <Chip
                label={evidence.status}
                size="small"
                color={evidence.status === 'APPROVED' ? 'success' : evidence.status === 'SUBMITTED' ? 'default' : 'warning'}
              />
            </Box>
          </ListItem>
          {index < evidences.length - 1 && <Divider />}
        </React.Fragment>
      ))}
    </List>
  );
};
