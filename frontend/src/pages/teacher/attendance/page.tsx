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
} from '@mui/material';
import {
  CalendarToday as CalendarIcon,
  CheckCircle as CheckCircleIcon,
  Cancel as CancelIcon,
  Save as SaveIcon,
} from '@mui/icons-material';

const TeacherAttendance = () => {
  const students = [
    { id: '1', name: 'John Doe', status: 'present' },
    { id: '2', name: 'Jane Smith', status: 'present' },
    { id: '3', name: 'Bob Johnson', status: 'absent' },
    { id: '4', name: 'Alice Williams', status: 'present' },
    { id: '5', name: 'Charlie Brown', status: 'late' },
    { id: '6', name: 'Diana Prince', status: 'present' },
    { id: '7', name: 'Edward King', status: 'absent' },
    { id: '8', name: 'Fiona Green', status: 'present' },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Attendance
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Record and manage student attendance for your classes
        </Typography>
      </Box>

      <Grid container spacing={3}>
        <Grid size={{ xs: 12, md: 4 }}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Select Class
              </Typography>
              <TextField
                select
                fullWidth
                label="Class"
                defaultValue="10-A"
                SelectProps={{
                  native: true,
                }}
              >
                <option value="10-A">Mathematics 10-A</option>
                <option value="10-B">Science 10-B</option>
                <option value="9-A">Mathematics 9-A</option>
              </TextField>
              <TextField
                fullWidth
                label="Date"
                type="date"
                defaultValue={new Date().toISOString().split('T')[0]}
                sx={{ mt: 2 }}
                InputLabelProps={{
                  shrink: true,
                }}
              />
              <Button
                variant="contained"
                fullWidth
                startIcon={<CalendarIcon />}
                sx={{ mt: 2 }}
              >
                Load Attendance
              </Button>
            </CardContent>
          </Card>

          <Card sx={{ mt: 2 }}>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Attendance Summary
              </Typography>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 1 }}>
                <Typography variant="body2">Present:</Typography>
                <Typography variant="body2" color="success.main">5</Typography>
              </Box>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 1 }}>
                <Typography variant="body2">Absent:</Typography>
                <Typography variant="body2" color="error.main">2</Typography>
              </Box>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 1 }}>
                <Typography variant="body2">Late:</Typography>
                <Typography variant="body2" color="warning.main">1</Typography>
              </Box>
              <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                <Typography variant="body2" fontWeight="medium">Total:</Typography>
                <Typography variant="body2" fontWeight="medium">8</Typography>
              </Box>
            </CardContent>
          </Card>
        </Grid>

        <Grid size={{ xs: 12, md: 8 }}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
                <Typography variant="h6">
                  Mathematics 10-A - {new Date().toLocaleDateString()}
                </Typography>
                <Button
                  variant="contained"
                  startIcon={<SaveIcon />}
                >
                  Save Attendance
                </Button>
              </Box>

              {students.map((student) => (
                <Box
                  key={student.id}
                  sx={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    py: 2,
                    borderBottom: '1px solid',
                    borderColor: 'divider',
                  }}
                >
                  <Typography variant="body1">{student.name}</Typography>
                  <Box sx={{ display: 'flex', gap: 1 }}>
                    <Chip
                      label="Present"
                      color={student.status === 'present' ? 'success' : 'default'}
                      onClick={() => {}}
                      clickable
                      icon={<CheckCircleIcon />}
                    />
                    <Chip
                      label="Absent"
                      color={student.status === 'absent' ? 'error' : 'default'}
                      onClick={() => {}}
                      clickable
                      icon={<CancelIcon />}
                    />
                    <Chip
                      label="Late"
                      color={student.status === 'late' ? 'warning' : 'default'}
                      onClick={() => {}}
                      clickable
                    />
                  </Box>
                </Box>
              ))}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
};

export default TeacherAttendance;
