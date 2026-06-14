/**
 * Assessment Form Component
 * Form for creating and editing assessments
 */

import { Box, TextField, Button, Typography, MenuItem } from '@mui/material';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { CreateAssessmentRequest, UpdateAssessmentRequest } from '@/shared/types/domain';

interface AssessmentFormProps {
  initialValues?: Partial<CreateAssessmentRequest>;
  onSubmit: (values: CreateAssessmentRequest | UpdateAssessmentRequest) => void;
  onCancel?: () => void;
  isEdit?: boolean;
}

const validationSchema = Yup.object({
  tp_id: Yup.string().required('TP is required'),
  tp_version_no: Yup.number().required('TP version is required').min(1),
  assessment_type: Yup.string().required('Assessment type is required'),
  assessment_items: Yup.array().required('Assessment items are required'),
  answer_key: Yup.array().required('Answer key is required'),
});

export const AssessmentForm = ({ initialValues, onSubmit, onCancel, isEdit = false }: AssessmentFormProps) => {
  const defaultValues: CreateAssessmentRequest = {
    tp_id: '',
    tp_version_no: 1,
    success_criteria_snapshot: {
      mastery_thresholds: [],
      performance_indicators: [],
      minimum_requirements: [],
    },
    assessment_type: 'FORMATIVE',
    assessment_items: {
      questions: [],
      total_score: 0,
      duration_minutes: 0,
    },
    answer_key: {
      version: '1.0',
      answers: {},
      notes: {},
    },
    scoring_guidelines: {
      version: '1.0',
      rubric: [],
      grading_scale: [],
    },
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        {isEdit ? 'Edit Assessment' : 'Create New Assessment'}
      </Typography>
      <Formik
        initialValues={{ ...defaultValues, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values: CreateAssessmentRequest, helpers: FormikHelpers<CreateAssessmentRequest>) => {
          onSubmit(values);
          helpers.resetForm();
        }}
      >
        {({ values, handleChange, touched, errors, isSubmitting }) => (
          <Form>
            <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
              <TextField
                fullWidth
                label="TP ID"
                name="tp_id"
                value={values.tp_id}
                onChange={handleChange}
                error={touched.tp_id && Boolean(errors.tp_id)}
                helperText={touched.tp_id && errors.tp_id}
              />

              <TextField
                fullWidth
                type="number"
                label="TP Version Number"
                name="tp_version_no"
                value={values.tp_version_no}
                onChange={handleChange}
                error={touched.tp_version_no && Boolean(errors.tp_version_no)}
                helperText={touched.tp_version_no && errors.tp_version_no}
              />

              <TextField
                select
                fullWidth
                label="Assessment Type"
                name="assessment_type"
                value={values.assessment_type}
                onChange={handleChange}
                error={touched.assessment_type && Boolean(errors.assessment_type)}
                helperText={touched.assessment_type && errors.assessment_type}
              >
                <MenuItem value="formative">Formative</MenuItem>
                <MenuItem value="summative">Summative</MenuItem>
                <MenuItem value="diagnostic">Diagnostic</MenuItem>
              </TextField>

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Assessment Items (JSON)"
                name="assessment_items"
                value={JSON.stringify(values.assessment_items, null, 2)}
                onChange={(e) => {
                  try {
                    const parsed = JSON.parse(e.target.value);
                    handleChange({ target: { name: 'assessment_items', value: parsed } });
                  } catch {
                    // Keep as string if invalid JSON
                  }
                }}
                error={touched.assessment_items && Boolean(errors.assessment_items)}
                helperText={touched.assessment_items && typeof errors.assessment_items === 'string' ? errors.assessment_items : undefined}
              />

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Answer Key (JSON)"
                name="answer_key"
                value={JSON.stringify(values.answer_key, null, 2)}
                onChange={(e) => {
                  try {
                    const parsed = JSON.parse(e.target.value);
                    handleChange({ target: { name: 'answer_key', value: parsed } });
                  } catch {
                    // Keep as string if invalid JSON
                  }
                }}
                error={touched.answer_key && Boolean(errors.answer_key)}
                helperText={touched.answer_key && typeof errors.answer_key === 'string' ? errors.answer_key : undefined}
              />

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Scoring Guidelines (JSON)"
                name="scoring_guidelines"
                value={JSON.stringify(values.scoring_guidelines, null, 2)}
                onChange={(e) => {
                  try {
                    const parsed = JSON.parse(e.target.value);
                    handleChange({ target: { name: 'scoring_guidelines', value: parsed } });
                  } catch {
                    // Keep as string if invalid JSON
                  }
                }}
              />

              <Box sx={{ display: 'flex', flexDirection: 'row', gap: 2, justifyContent: 'flex-end' }}>
                {onCancel && (
                  <Button variant="outlined" onClick={onCancel} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
                <Button type="submit" variant="contained" disabled={isSubmitting}>
                  {isSubmitting ? 'Saving...' : isEdit ? 'Update Assessment' : 'Create Assessment'}
                </Button>
              </Box>
            </Box>
          </Form>
        )}
      </Formik>
    </Box>
  );
};
