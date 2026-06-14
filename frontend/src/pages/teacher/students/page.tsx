import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
  Chip,
  TextField,
  Avatar,
} from '@mui/material';
import {
  Search as SearchIcon,
  TrendingUp as TrendingUpIcon,
  Visibility as VisibilityIcon,
} from '@mui/icons-material';

const TeacherStudents = () => {
  const students = [
    { id: '1', name: 'John Doe', class: '10-A', progress: 85, attendance: 92 },
    { id: '2', name: 'Jane Smith', class: '10-A', progress: 78, attendance: 95 },
    { id: '3', name: 'Bob Johnson', class: '10-A', progress: 65, attendance: 88 },
    { id: '4', name: 'Alice Williams', class: '10-A', progress: 92, attendance: 98 },
    { id: '5', name: 'Charlie Brown', class: '10-A', progress: 70, attendance: 85 },
    { id: '6', name: 'Diana Prince', class: '10-A', progress: 88, attendance: 90 },
    { id: '7', name: 'Edward King', class: '10-A', progress: 75, attendance: 82 },
    { id: '8', name: 'Fiona Green', class: '10-A', progress: 80, attendance: 94 },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          My Students
        </Typography>
        <Typography variant="body2" color="text.secondary">
          View and manage student progress across your classes
        </Typography>
      </Box>

      <Card>
        <CardContent>
          <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
            <TextField
              placeholder="Search students..."
              size="small"
              InputProps={{
                startAdornment: <SearchIcon sx={{ mr: 1, color: 'text.secondary' }} />,
              }}
              sx={{ width: 300 }}
            />
            <Box sx={{ display: 'flex', gap: 2 }}>
              <Button variant="outlined">Filter by Class</Button>
              <Button variant="outlined">Export List</Button>
            </Box>
          </Box>

          <Grid container spacing={3}>
            {students.map((student) => (
              <Grid size={{ xs: 12, sm: 6, md: 4, lg: 3 }} key={student.id}>
                <Card variant="outlined">
                  <CardContent>
                    <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
                      <Avatar sx={{ mr: 2 }}>{student.name.charAt(0)}</Avatar>
                      <Box sx={{ flexGrow: 1 }}>
                        <Typography variant="subtitle1" fontWeight="medium">
                          {student.name}
                        </Typography>
                        <Typography variant="body2" color="text.secondary">
                          Class {student.class}
                        </Typography>
                      </Box>
                    </Box>

                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, mb: 1 }}>
                      <TrendingUpIcon fontSize="small" sx={{ color: 'text.secondary' }} />
                      <Typography variant="body2" color="text.secondary">
                        Progress: {student.progress}%
                      </Typography>
                    </Box>
                    <Box
                      sx={{
                        width: '100%',
                        height: 6,
                        backgroundColor: 'grey.200',
                        borderRadius: 3,
                        overflow: 'hidden',
                        mb: 2,
                      }}
                    >
                      <Box
                        sx={{
                          width: `${student.progress}%`,
                          height: '100%',
                          backgroundColor: student.progress >= 80 ? 'success.main' : student.progress >= 60 ? 'warning.main' : 'error.main',
                          borderRadius: 3,
                        }}
                      />
                    </Box>

                    <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mt: 2 }}>
                      <Chip
                        label={`Attendance: ${student.attendance}%`}
                        size="small"
                        color={student.attendance >= 90 ? 'success' : 'default'}
                      />
                      <Button size="small" startIcon={<VisibilityIcon />}>
                        View
                      </Button>
                    </Box>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>
        </CardContent>
      </Card>
    </Container>
  );
};

export default TeacherStudents;
