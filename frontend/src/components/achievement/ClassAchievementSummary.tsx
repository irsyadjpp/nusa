/**
 * Class Achievement Summary Component
 * Summary for entire class
 */

import React from 'react';
import {
  Box,
  Typography,
  Card,
  CardContent,
  Grid,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Chip,
  Avatar,
} from '@mui/material';
import {
  TrendingUp,
  People,
  Star,
  Warning,
} from '@mui/icons-material';

interface MasteryDistribution {
  excellent: number;
  proficient: number;
  developing: number;
  beginning: number;
}

interface CompetencyAchievement {
  competency_id: string;
  competency_name: string;
  average_score: number;
  mastery_distribution: MasteryDistribution;
}

interface TopPerformer {
  student_id: string;
  student_name: string;
  average_score: number;
}

interface AreaForImprovement {
  competency_id: string;
  competency_name: string;
  average_score: number;
  struggling_students: number;
}

interface ClassAchievement {
  class_id: string;
  class_name: string;
  subject_id: string;
  subject_name: string;
  total_students: number;
  average_mastery: number;
  competency_achievements: CompetencyAchievement[];
  top_performers: TopPerformer[];
  areas_for_improvement: AreaForImprovement[];
}

interface ClassAchievementSummaryProps {
  data: ClassAchievement;
}

const ClassAchievementSummary: React.FC<ClassAchievementSummaryProps> = ({ data }) => {
  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        Ringkasan Pencapaian Kelas
      </Typography>

      <Grid container spacing={3}>
        {/* Overview Cards */}
        <Grid item xs={12} sm={6} md={3}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                <People color="primary" />
                <Typography variant="body2" color="text.secondary">
                  Total Siswa
                </Typography>
              </Box>
              <Typography variant="h4" fontWeight="bold">
                {data.total_students}
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} sm={6} md={3}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                <TrendingUp color="success" />
                <Typography variant="body2" color="text.secondary">
                  Rata-rata Penguasaan
                </Typography>
              </Box>
              <Typography variant="h4" fontWeight="bold">
                {data.average_mastery.toFixed(1)}%
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} sm={6} md={3}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                <Star color="warning" />
                <Typography variant="body2" color="text.secondary">
                  Top Performers
                </Typography>
              </Box>
              <Typography variant="h4" fontWeight="bold">
                {data.top_performers.length}
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} sm={6} md={3}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                <Warning color="error" />
                <Typography variant="body2" color="text.secondary">
                  Area Perlu Perbaikan
                </Typography>
              </Box>
              <Typography variant="h4" fontWeight="bold">
                {data.areas_for_improvement.length}
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        {/* Competency Achievements */}
        <Grid item xs={12}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Pencapaian Kompetensi
              </Typography>
              <TableContainer>
                <Table>
                  <TableHead>
                    <TableRow>
                      <TableCell>Kompetensi</TableCell>
                      <TableCell align="right">Rata-rata Skor</TableCell>
                      <TableCell align="right">Sangat Baik</TableCell>
                      <TableCell align="right">Baik</TableCell>
                      <TableCell align="right">Sedang Berkembang</TableCell>
                      <TableCell align="right">Perlu Bimbingan</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {data.competency_achievements.map((competency) => (
                      <TableRow key={competency.competency_id}>
                        <TableCell>{competency.competency_name}</TableCell>
                        <TableCell align="right">{competency.average_score.toFixed(1)}</TableCell>
                        <TableCell align="right">{competency.mastery_distribution.excellent}</TableCell>
                        <TableCell align="right">{competency.mastery_distribution.proficient}</TableCell>
                        <TableCell align="right">{competency.mastery_distribution.developing}</TableCell>
                        <TableCell align="right">{competency.mastery_distribution.beginning}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </CardContent>
          </Card>
        </Grid>

        {/* Top Performers */}
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Top Performers
              </Typography>
              {data.top_performers.map((performer, index) => (
                <Box
                  key={performer.student_id}
                  sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 2 }}
                >
                  <Avatar>{index + 1}</Avatar>
                  <Box sx={{ flexGrow: 1 }}>
                    <Typography variant="body1" fontWeight="medium">
                      {performer.student_name}
                    </Typography>
                    <Typography variant="caption" color="text.secondary">
                      Rata-rata: {performer.average_score.toFixed(1)}
                    </Typography>
                  </Box>
                  <Chip
                    label={`${performer.average_score.toFixed(1)}`}
                    color="success"
                    size="small"
                  />
                </Box>
              ))}
            </CardContent>
          </Card>
        </Grid>

        {/* Areas for Improvement */}
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Area Perlu Perbaikan
              </Typography>
              {data.areas_for_improvement.map((area) => (
                <Box
                  key={area.competency_id}
                  sx={{ p: 2, border: 1, borderColor: 'divider', borderRadius: 1, mb: 2 }}
                >
                  <Typography variant="body1" fontWeight="medium" gutterBottom>
                    {area.competency_name}
                  </Typography>
                  <Box sx={{ display: 'flex', justifyContent: 'space-between' }}>
                    <Typography variant="caption" color="text.secondary">
                      Rata-rata: {area.average_score.toFixed(1)}
                    </Typography>
                    <Typography variant="caption" color="error">
                      {area.struggling_students} siswa kesulitan
                    </Typography>
                  </Box>
                </Box>
              ))}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Box>
  );
};

export default ClassAchievementSummary;
