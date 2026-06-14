/**
 * Subject Category List Component
 * Displays a list of subject categories with hierarchy indicators
 */

import { Box, List, ListItem, ListItemText, ListItemButton, Chip, Typography } from '@mui/material';
import { SubjectCategory } from '@/api/academic-foundation';

interface SubjectCategoryListProps {
  subjectCategories: SubjectCategory[];
  selectedId?: string;
  onSelect?: (subjectCategory: SubjectCategory) => void;
  loading?: boolean;
}

export const SubjectCategoryList = ({ subjectCategories, selectedId, onSelect, loading }: SubjectCategoryListProps) => {
  if (loading) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading subject categories...</Typography>
      </Box>
    );
  }

  if (subjectCategories.length === 0) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography color="text.secondary">No subject categories found</Typography>
      </Box>
    );
  }

  return (
    <List>
      {subjectCategories.map((category) => (
        <ListItem
          key={category.id}
          disablePadding
          sx={{
            backgroundColor: selectedId === category.id ? 'action.selected' : 'transparent',
            pl: category.level * 2,
          }}
        >
          <ListItemButton onClick={() => onSelect?.(category)}>
            <ListItemText
              primary={`${category.code} - ${category.name_indonesian}`}
              secondary={`Level: ${category.level} | Order: ${category.sort_order}`}
            />
          </ListItemButton>
          <Box sx={{ display: 'flex', alignItems: 'center', pr: 2 }}>
            <Chip
              label={category.is_active ? 'Active' : 'Inactive'}
              size="small"
              color={category.is_active ? 'success' : 'default'}
            />
          </Box>
        </ListItem>
      ))}
    </List>
  );
};