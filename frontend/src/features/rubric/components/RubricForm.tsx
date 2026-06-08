/**
 * Rubric Form Component
 * Form for creating and editing rubrics
 */

import { Box, TextField, Button, Stack, Typography } from '@mui/material';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';

interface RubricFormProps {
  initialValues?: any;
  onSubmit: (values: any) => void;
  onCancel?: () => void;
  isEdit?: boolean;
}

const validationSchema = Yup.object({
  title: Yup.string().required('Title is required'),
  assessment_id: Yup.string().required('Assessment is required'),
  criteria: Yup.array().required('Criteria are required'),
});

export const RubricForm = ({ initialValues, onSubmit, onCancel, isEdit = false }: RubricFormProps) => {
  const defaultValues = {
    title: '',
    assessment_id: '',
    description: '',
    criteria: [],
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        {isEdit ? 'Edit Rubric' : 'Create New Rubric'}
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
                label="Title"
                name="title"
                value={values.title}
                onChange={handleChange}
                error={touched.title && Boolean(errors.title)}
                helperText={touched.title && errors.title}
              />

              <TextField
                fullWidth
                label="Assessment ID"
                name="assessment_id"
                value={values.assessment_id}
                onChange={handleChange}
                error={touched.assessment_id && Boolean(errors.assessment_id)}
                helperText={touched.assessment_id && errors.assessment_id}
              />

              <TextField
                fullWidth
                multiline
                rows={3}
                label="Description"
                name="description"
                value={values.description}
                onChange={handleChange}
              />

              <TextField
                fullWidth
                multiline
                rows={6}
                label="Criteria (JSON)"
                name="criteria"
                value={JSON.stringify(values.criteria, null, 2)}
                onChange={(e) => {
                  try {
                    const parsed = JSON.parse(e.target.value);
                    handleChange({ target: { name: 'criteria', value: parsed } });
                  } catch {
                    // Keep as string if invalid JSON
                  }
                }}
                error={touched.criteria && Boolean(errors.criteria)}
                helperText={touched.criteria && errors.criteria}
              />

              <Stack direction="row" spacing={2} justifyContent="flex-end">
                {onCancel && (
                  <Button variant="outlined" onClick={onCancel} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
                <Button type="submit" variant="contained" disabled={isSubmitting}>
                  {isSubmitting ? 'Saving...' : isEdit ? 'Update Rubric' : 'Create Rubric'}
                </Button>
              </Stack>
            </Stack>
          </Form>
        )}
      </Formik>
    </Box>
  );
};
