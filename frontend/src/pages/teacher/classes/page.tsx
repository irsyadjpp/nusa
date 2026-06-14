import {
  Box,
  Button,
  Card,
  CardContent,
  Container,
  Typography,
  Grid,
  Chip,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  People as PeopleIcon,
  Assignment as AssignmentIcon,
  ArrowForward as ArrowForwardIcon,
} from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { useClasses } from '@/services/queries/ClassesQueryService';

const TeacherClasses = () => {
  const navigate = useNavigate();
  const { data: classesData, isLoading, error } = useClasses();

  if (isLoading) {
    return (
      <Container maxWidth="xl">
        <Box sx={{ display: 'flex', justifyContent: 'center', mt: 10 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="xl">
        <Alert severity="error">{error.message}</Alert>
      </Container>
    );
  }

  const classes = classesData?.data || [
    {
      id: '1',
      name: 'Mathematics 10-A',
      subject: 'Mathematics',
      grade: '10',
      section: 'A',
      students: 32,
      schedule: 'Mon, Wed, Fri - 08:00-09:30',
      room: 'Room 101',
      progress: 75,
    },
    {
      id: '2',
      name: 'Science 10-B',
      subject: 'Science',
      grade: '10',
      section: 'B',
      students: 28,
      schedule: 'Tue, Thu - 10:00-11:30',
      room: 'Room 205',
      progress: 60,
    },
    {
      id: '3',
      name: 'Mathematics 9-A',
      subject: 'Mathematics',
      grade: '9',
      section: 'A',
      students: 35,
      schedule: 'Mon, Wed, Fri - 13:00-14:30',
      room: 'Room 101',
      progress: 45,
    },
    {
      id: '4',
      name: 'Physics 11-A',
      subject: 'Physics',
      grade: '11',
      section: 'A',
      students: 25,
      schedule: 'Tue, Thu - 14:00-15:30',
      room: 'Lab 301',
      progress: 30,
    },
  ];

  return (
    <Container maxWidth="xl">
      <Box sx={{ mb: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          My Classes
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Manage your assigned classes and track student progress
        </Typography>
      </Box>

      <Grid container spacing={3}>
        {classes.map((classItem: any) => (
          <Grid size={{ xs: 12, md: 6, lg: 4 }} key={classItem.id}>
            <Card sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
              <CardContent sx={{ flexGrow: 1 }}>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', mb: 2 }}>
                  <Typography variant="h6" gutterBottom>
                    {classItem.name}
                  </Typography>
                  <Chip label={classItem.subject} size="small" color="primary" />
                </Box>

                <Box sx={{ display: 'flex', alignItems: 'center', gap: 3, mb: 2 }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
                    <Box sx={{ color: 'text.secondary' }}>
                      <PeopleIcon fontSize="small" />
                    </Box>
                    <Typography variant="body2" color="text.secondary">
                      {classItem.students} Students
                    </Typography>
                  </Box>
                  <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
                    <Box sx={{ color: 'text.secondary' }}>
                      <AssignmentIcon fontSize="small" />
                    </Box>
                    <Typography variant="body2" color="text.secondary">
                      {classItem.room}
                    </Typography>
                  </Box>
                </Box>

                <Typography variant="body2" color="text.secondary" gutterBottom>
                  {classItem.schedule}
                </Typography>

                <Box sx={{ mt: 2 }}>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 0.5 }}>
                    <Typography variant="body2" color="text.secondary">
                      Curriculum Progress
                    </Typography>
                    <Typography variant="body2" fontWeight="medium">
                      {classItem.progress}%
                    </Typography>
                  </Box>
                  <Box
                    sx={{
                      width: '100%',
                      height: 6,
                      backgroundColor: 'grey.200',
                      borderRadius: 3,
                      overflow: 'hidden',
                    }}
                  >
                    <Box
                      sx={{
                        width: `${classItem.progress}%`,
                        height: '100%',
                        backgroundColor: 'primary.main',
                        borderRadius: 3,
                      }}
                    />
                  </Box>
                </Box>
              </CardContent>

              <Box sx={{ p: 2, pt: 0 }}>
                <Button
                  variant="outlined"
                  fullWidth
                  endIcon={<ArrowForwardIcon />}
                  onClick={() => navigate(`/teacher/classes/${classItem.id}`)}
                >
                  View Class Details
                </Button>
              </Box>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default TeacherClasses;
