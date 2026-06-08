/**
 * Evidence List Component
 * Displays a list of evidences with filtering and pagination
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { Evidence } from '@/api/evidence';

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
      {evidences.map((evidence) => (
        <ListItem
          key={evidence.id}
          disablePadding
          selected={selectedId === evidence.id}
          secondaryAction={
            <Chip
              label={evidence.status}
              size="small"
              color={evidence.status === 'approved' ? 'success' : evidence.status === 'draft' ? 'default' : 'warning'}
            />
          }
        >
          <ListItemButton onClick={() => onSelect?.(evidence)} selected={selectedId === evidence.id}>
            <ListItemText
              primary={evidence.title || `Evidence ${evidence.id}`}
              secondary={`Student: ${evidence.student_id} | Assessment: ${evidence.assessment_id}`}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};
