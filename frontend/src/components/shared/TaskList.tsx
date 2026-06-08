/**
 * TaskList Component
 * Displays a list of tasks with checkboxes
 */

import { List, ListItem, ListItemButton, ListItemText, Checkbox, Typography } from '@mui/material';

interface Task {
  id: string;
  title: string;
  completed: boolean;
  dueDate?: string;
}

interface TaskListProps {
  tasks: Task[];
  onToggle?: (taskId: string) => void;
  showDueDate?: boolean;
}

export const TaskList = ({ tasks, onToggle, showDueDate = false }: TaskListProps) => {
  return (
    <List>
      {tasks.map((task) => (
        <ListItem
          key={task.id}
          disablePadding
          secondaryAction={showDueDate && task.dueDate && (
            <Typography variant="caption" color="text.secondary">
              {task.dueDate}
            </Typography>
          )}
        >
          <ListItemButton onClick={() => onToggle?.(task.id)} dense>
            <Checkbox
              edge="start"
              checked={task.completed}
              tabIndex={-1}
              disableRipple
            />
            <ListItemText
              primary={task.title}
              sx={{
                textDecoration: task.completed ? 'line-through' : 'none',
                color: task.completed ? 'text.secondary' : 'text.primary',
              }}
            />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  );
};
