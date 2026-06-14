import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import {
  Box,
  Typography,
  Button,
  Paper,
  Grid,
  CircularProgress,
  Alert,
  Card,
  CardContent,
  Chip,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  LinearProgress,
  Container,
  Breadcrumbs,
  Link,
  TextField,
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useNavigate } from 'react-router-dom';
import { useClassAchievement } from '@/services/queries/AchievementQueryService';

const ClassDashboard: React.FC = () => {
  const navigate = useNavigate();
  const { id: classId } = useParams<{ id: string }>();
  const [subjectId, setSubjectId] = useState('');

  const {
    data: classAchievement,
    isLoading,
    error,
    refetch,
  } = useClassAchievement(classId || '', { subject_id: subjectId || undefined });

  const handleFilter = () => {
    refetch();
  };

  if (isLoading) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'center', py: 8 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="xl" sx={{ py: 4 }}>
        <Alert severity="error">{error.message || 'Failed to load class data'}</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="xl" sx={{ py: 4 }}>
      <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
        <Link underline="hover" onClick={() => navigate('/achievement')}>
          Achievement
        </Link>
        <Typography color="text.primary">Class Dashboard</Typography>
      </Breadcrumbs>

      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
        <Typography variant="h4" component="h1">
          Class Dashboard - {classId}
        </Typography>
        <Button
          variant="outlined"
          startIcon={<ArrowBackIcon />}
          onClick={() => navigate('/achievement')}
        >
          Back to Achievement
        </Button>
      </Box>

      <Paper sx={{ p: 2, mb: 3 }}>
        <Grid container spacing={2} alignItems="center">
          <Grid size={{ xs: 12, md: 4 }}>
            <TextField
              label="Subject ID"
              value={subjectId}
              onChange={(e) => setSubjectId(e.target.value)}
              fullWidth
              size="small"
            />
          </Grid>
          <Grid size={{ xs: 12, md: 8 }}>
            <Button variant="contained" onClick={handleFilter}>
              Apply Filter
            </Button>
          </Grid>
        </Grid>
      </Paper>

      {classAchievement ? (
        <Grid container spacing={3}>
          <Grid size={{ xs: 12, md: 4 }}>
            <Card>
              <CardContent>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Total Students
                </Typography>
                <Typography variant="h4">{classAchievement.total_students}</Typography>
              </CardContent>
            </Card>
          </Grid>

          <Grid size={{ xs: 12, md: 4 }}>
            <Card>
              <CardContent>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Average Mastery
                </Typography>
                <Typography variant="h4">{classAchievement.average_mastery}%</Typography>
                <LinearProgress
                  variant="determinate"
                  value={classAchievement.average_mastery}
                  sx={{ mt: 1 }}
                />
              </CardContent>
            </Card>
          </Grid>

          <Grid size={{ xs: 12, md: 4 }}>
            <Card>
              <CardContent>
                <Typography variant="body2" color="text.secondary" gutterBottom>
                  Overall Class Score
                </Typography>
                <Typography variant="h4">{classAchievement.overall_class_score}</Typography>
              </CardContent>
            </Card>
          </Grid>

          {classAchievement.competency_achievements && classAchievement.competency_achievements.length > 0 && (
            <Grid size={{ xs: 12 }}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    Competency Achievements
                  </Typography>
                  <TableContainer>
                    <Table>
                      <TableHead>
                        <TableRow>
                          <TableCell>Competency</TableCell>
                          <TableCell>Average Score</TableCell>
                          <TableCell>Mastery Level</TableCell>
                          <TableCell>Mastery Distribution</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {classAchievement.competency_achievements.map((comp: any) => (
                          <TableRow key={comp.competency_id}>
                            <TableCell>{comp.competency_name}</TableCell>
                            <TableCell>{comp.average_score}</TableCell>
                            <TableCell>
                              <Chip
                                label={comp.mastery_level}
                                color={
                                  comp.mastery_level === 'EXEMPLARY'
                                    ? 'success'
                                    : comp.mastery_level === 'PROFICIENT'
                                    ? 'primary'
                                    : comp.mastery_level === 'DEVELOPING'
                                    ? 'warning'
                                    : 'error'
                                }
                                size="small"
                              />
                            </TableCell>
                            <TableCell>
                              <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap' }}>
                                <Chip label={`Excellent: ${comp.mastery_distribution.excellent}`} size="small" color="success" />
                                <Chip label={`Proficient: ${comp.mastery_distribution.proficient}`} size="small" color="primary" />
                                <Chip label={`Developing: ${comp.mastery_distribution.developing}`} size="small" color="warning" />
                                <Chip label={`Beginning: ${comp.mastery_distribution.beginning}`} size="small" color="error" />
                              </Box>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </CardContent>
              </Card>
            </Grid>
          )}

          {classAchievement.student_achievements && classAchievement.student_achievements.length > 0 && (
            <Grid size={{ xs: 12 }}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    Student Achievements
                  </Typography>
                  <TableContainer>
                    <Table>
                      <TableHead>
                        <TableRow>
                          <TableCell>Student</TableCell>
                          <TableCell>Average Score</TableCell>
                          <TableCell>Mastery Level</TableCell>
                          <TableCell>Completed Competencies</TableCell>
                          <TableCell>Action</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {classAchievement.student_achievements.map((student: any) => (
                          <TableRow key={student.student_id}>
                            <TableCell>{student.student_name}</TableCell>
                            <TableCell>{student.average_score}</TableCell>
                            <TableCell>
                              <Chip
                                label={student.mastery_level}
                                color={
                                  student.mastery_level === 'EXEMPLARY'
                                    ? 'success'
                                    : student.mastery_level === 'PROFICIENT'
                                    ? 'primary'
                                    : student.mastery_level === 'DEVELOPING'
                                    ? 'warning'
                                    : 'error'
                                }
                                size="small"
                              />
                            </TableCell>
                            <TableCell>
                              {student.completed_competencies} / {student.total_competencies}
                            </TableCell>
                            <TableCell>
                              <Button
                                size="small"
                                onClick={() => navigate(`/dashboard/student/${student.student_id}`)}
                              >
                                View Details
                              </Button>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </CardContent>
              </Card>
            </Grid>
          )}

          {classAchievement.areas_for_improvement && classAchievement.areas_for_improvement.length > 0 && (
            <Grid size={{ xs: 12 }}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    Areas for Improvement
                  </Typography>
                  <Grid container spacing={2}>
                    {classAchievement.areas_for_improvement.map((area: any) => (
                      <Grid size={{ xs: 12, md: 6 }} key={area.competency_id}>
                        <Paper sx={{ p: 2 }}>
                          <Typography variant="subtitle2" gutterBottom>
                            {area.competency_name}
                          </Typography>
                          <Typography variant="body2" color="text.secondary">
                            Average Score: {area.average_score}
                          </Typography>
                          <Typography variant="body2" color="text.secondary">
                            Struggling Students: {area.struggling_students}
                          </Typography>
                        </Paper>
                      </Grid>
                    ))}
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
          )}

          {classAchievement.top_performers && classAchievement.top_performers.length > 0 && (
            <Grid size={{ xs: 12 }}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    Top Performers
                  </Typography>
                  <TableContainer>
                    <Table>
                      <TableHead>
                        <TableRow>
                          <TableCell>Student</TableCell>
                          <TableCell>Average Score</TableCell>
                          <TableCell>Mastery Level</TableCell>
                          <TableCell>Action</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {classAchievement.top_performers.map((student: any) => (
                          <TableRow key={student.student_id}>
                            <TableCell>{student.student_name}</TableCell>
                            <TableCell>{student.average_score}</TableCell>
                            <TableCell>
                              <Chip
                                label={student.mastery_level}
                                color={
                                  student.mastery_level === 'EXEMPLARY'
                                    ? 'success'
                                    : student.mastery_level === 'PROFICIENT'
                                    ? 'primary'
                                    : student.mastery_level === 'DEVELOPING'
                                    ? 'warning'
                                    : 'error'
                                }
                                size="small"
                              />
                            </TableCell>
                            <TableCell>
                              <Button
                                size="small"
                                onClick={() => navigate(`/dashboard/student/${student.student_id}`)}
                              >
                                View Details
                              </Button>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </CardContent>
              </Card>
            </Grid>
          )}

          <Grid size={{ xs: 12 }}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Actions
                </Typography>
                <Box sx={{ display: 'flex', gap: 2 }}>
                  <Button variant="contained" onClick={() => navigate('/reports/generate', { state: { studentId: null, classId } })}>
                    Generate Class Narrative Report
                  </Button>
                  <Button variant="outlined" onClick={() => navigate('/achievement')}>
                    Back to Achievement Dashboard
                  </Button>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      ) : (
        <Alert severity="info">No class achievement data available</Alert>
      )}
    </Container>
  );
};

export default ClassDashboard;
