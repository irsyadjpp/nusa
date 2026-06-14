/**
 * Evaluation Form Component
 * Form for creating and editing evaluations
 */

import { Box, TextField, Button, Stack, Typography } from '@mui/material';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';

interface EvaluationFormProps {
  initialValues?: any;
  onSubmit: (values: any) => void;
  onCancel?: () => void;
  isEdit?: boolean;
}

const validationSchema = Yup.object({
  evidence_id: Yup.string().required('Evidence is required'),
  evaluator_id: Yup.string().required('Evaluator is required'),
  score: Yup.number().required('Score is required').min(0).max(100),
  feedback: Yup.string().required('Feedback is required'),
});

export const EvaluationForm = ({ initialValues, onSubmit, onCancel, isEdit = false }: EvaluationFormProps) => {
  const defaultValues = {
    evidence_id: '',
    evaluator_id: '',
    score: 0,
    feedback: '',
    rubric_scores: [],
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        {isEdit ? 'Edit Evaluation' : 'Create New Evaluation'}
      </Typography>
      <Formik
        initialValues={{ ...defaultValues, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values: any, helpers: FormikHelpers<any>) => {
          onSubmit(values);
          helpers.resetForm();
        }}
      >
        {({ values, handleChange, touched, errors, isSubmitting }) => (
          <Form>
            <Stack spacing={3}>
              <TextField
                fullWidth
                label="Evidence ID"
                name="evidence_id"
                value={values.evidence_id}
                onChange={handleChange}
                error={touched.evidence_id && Boolean(errors.evidence_id)}
                helperText={touched.evidence_id && typeof errors.evidence_id === 'string' ? errors.evidence_id : undefined}
              />

              <TextField
                fullWidth
                label="Evaluator ID"
                name="evaluator_id"
                value={values.evaluator_id}
                onChange={handleChange}
                error={touched.evaluator_id && Boolean(errors.evaluator_id)}
                helperText={touched.evaluator_id && typeof errors.evaluator_id === 'string' ? errors.evaluator_id : undefined}
              />

              <TextField
                fullWidth
                type="number"
                label="Score"
                name="score"
                value={values.score}
                onChange={handleChange}
                error={touched.score && Boolean(errors.score)}
                helperText={touched.score && typeof errors.score === 'string' ? errors.score : undefined}
                inputProps={{ min: 0, max: 100 }}
              />

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Feedback"
                name="feedback"
                value={values.feedback}
                onChange={handleChange}
                error={touched.feedback && Boolean(errors.feedback)}
                helperText={touched.feedback && typeof errors.feedback === 'string' ? errors.feedback : undefined}
              />

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Rubric Scores (JSON)"
                name="rubric_scores"
                value={JSON.stringify(values.rubric_scores, null, 2)}
                onChange={(e) => {
                  try {
                    const parsed = JSON.parse(e.target.value);
                    handleChange({ target: { name: 'rubric_scores', value: parsed } });
                  } catch {
                    // Keep as string if invalid JSON
                  }
                }}
              />

              <Stack direction="row" spacing={2} justifyContent="flex-end">
                {onCancel && (
                  <Button variant="outlined" onClick={onCancel} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
                <Button type="submit" variant="contained" disabled={isSubmitting}>
                  {isSubmitting ? 'Saving...' : isEdit ? 'Update Evaluation' : 'Create Evaluation'}
                </Button>
              </Stack>
            </Stack>
          </Form>
        )}
      </Formik>
    </Box>
  );
};
