package expression

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/expression/interpolater"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/pkg"
	"path/filepath"
)

var _ = Context("EvalToFile", func() {
	Context("expression is float64", func() {
		It("should call data.CoerceToFile w/ expected args", func() {
			/* arrange */
			providedExpression := 2.2
			providedScratchDir := "dummyScratchDir"

			fakeData := new(data.Fake)

			objectUnderTest := _fileEvaluator{
				data: fakeData,
			}

			/* act */
			objectUnderTest.EvalToFile(
				map[string]*model.Value{},
				providedExpression,
				new(pkg.FakeHandle),
				providedScratchDir,
			)

			/* assert */
			actualValue,
				actualScratchDir := fakeData.CoerceToFileArgsForCall(0)
			Expect(*actualValue).To(Equal(model.Value{Number: &providedExpression}))
			Expect(actualScratchDir).To(Equal(providedScratchDir))
		})
		It("should return expected result", func() {
			/* arrange */
			fakeData := new(data.Fake)
			coercedValue := model.Value{Number: new(float64)}
			coerceToFileErr := errors.New("dummyError")

			fakeData.CoerceToFileReturns(&coercedValue, coerceToFileErr)

			objectUnderTest := _fileEvaluator{
				data: fakeData,
			}

			/* act */
			actualValue, actualErr := objectUnderTest.EvalToFile(
				map[string]*model.Value{},
				2.2,
				new(pkg.FakeHandle),
				"dummyScratchDir",
			)

			/* assert */
			Expect(*actualValue).To(Equal(coercedValue))
			Expect(actualErr).To(Equal(coerceToFileErr))
		})
	})
	Context("expression is map[string]interface{}", func() {
		It("should call data.CoerceToFile w/ expected args", func() {
			/* arrange */
			providedExpression := map[string]interface{}{"dummyName": 2.2}
			providedScratchDir := "dummyScratchDir"

			fakeData := new(data.Fake)

			objectUnderTest := _fileEvaluator{
				data: fakeData,
			}

			/* act */
			objectUnderTest.EvalToFile(
				map[string]*model.Value{},
				providedExpression,
				new(pkg.FakeHandle),
				providedScratchDir,
			)

			/* assert */
			actualValue,
				actualScratchDir := fakeData.CoerceToFileArgsForCall(0)
			Expect(*actualValue).To(Equal(model.Value{Object: providedExpression}))
			Expect(actualScratchDir).To(Equal(providedScratchDir))
		})
		It("should return expected result", func() {
			/* arrange */
			fakeData := new(data.Fake)
			coercedValue := model.Value{Object: map[string]interface{}{}}
			coerceToFileErr := errors.New("dummyError")

			fakeData.CoerceToFileReturns(&coercedValue, coerceToFileErr)

			objectUnderTest := _fileEvaluator{
				data: fakeData,
			}

			/* act */
			actualValue, actualErr := objectUnderTest.EvalToFile(
				map[string]*model.Value{},
				map[string]interface{}{},
				new(pkg.FakeHandle),
				"dummyScratchDir",
			)

			/* assert */
			Expect(*actualValue).To(Equal(coercedValue))
			Expect(actualErr).To(Equal(coerceToFileErr))
		})
	})
	Context("expression is string", func() {
		Context("expression is pkg fs ref", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				providedScope := map[string]*model.Value{"dummyName": {}}

				pkgFSRef := "/dummyPkgFSRef"
				providedExpression := fmt.Sprintf("$(%v)", pkgFSRef)
				providedPkgRef := new(pkg.FakeHandle)

				fakeInterpolater := new(interpolater.Fake)
				// err to trigger immediate return
				fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

				objectUnderTest := _fileEvaluator{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.EvalToFile(
					providedScope,
					providedExpression,
					providedPkgRef,
					"dummyScratchDir",
				)

				/* assert */
				actualExpression,
					actualScope,
					actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

				Expect(actualExpression).To(Equal(pkgFSRef))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
			})
			Context("interpolater.Interpolate errs", func() {
				It("should return expected result", func() {

					/* arrange */
					pkgFSRef := "/dummyPkgFSRef"
					providedExpression := fmt.Sprintf("$(%v)", pkgFSRef)

					fakeInterpolater := new(interpolater.Fake)
					interpolateErr := errors.New("dummyError")
					fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

					expectedErr := fmt.Errorf(
						"unable to evaluate pkg fs ref %v; error was %v",
						pkgFSRef,
						interpolateErr.Error(),
					)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						providedExpression,
						new(pkg.FakeHandle),
						"dummyScratchDir",
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(expectedErr))

				})
			})
			Context("interpolater.Interpolate doesn't error", func() {
				It("should return expected result", func() {
					/* arrange */
					interpolatedPath := "dummyInterpolatedPath"
					fakeInterpolater := new(interpolater.Fake)
					fakeInterpolater.InterpolateReturns(interpolatedPath, nil)

					fakePkgHandle := new(pkg.FakeHandle)
					pkgHandleRef := "dummyPkgHandleRef"
					fakePkgHandle.RefReturns(pkgHandleRef)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					expectedValue := filepath.Join(pkgHandleRef, interpolatedPath)

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						"$(/dummyPkgFSRef)",
						fakePkgHandle,
						"dummyScratchDir",
					)

					/* assert */
					Expect(*actualValue).To(Equal(model.Value{File: &expectedValue}))
					Expect(actualErr).To(BeNil())
				})
			})
		})
		Context("expression is scope ref", func() {
			It("should call data.CoerceToFile w/ expected args", func() {
				/* arrange */
				scopeName := "dummyScopeName"
				providedExpression := fmt.Sprintf("$(%v)", scopeName)
				scopeValue := model.Value{File: new(string)}
				providedScope := map[string]*model.Value{
					scopeName: &scopeValue,
				}
				providedScratchDir := "dummyScratchDir"

				fakeData := new(data.Fake)

				objectUnderTest := _fileEvaluator{
					data: fakeData,
				}

				/* act */
				objectUnderTest.EvalToFile(
					providedScope,
					providedExpression,
					new(pkg.FakeHandle),
					providedScratchDir,
				)

				/* assert */
				actualValue, actualScratchDir := fakeData.CoerceToFileArgsForCall(0)
				Expect(*actualValue).To(Equal(scopeValue))
				Expect(actualScratchDir).To(Equal(providedScratchDir))
			})
		})
		Context("expression is deprecated pkg fs ref", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				providedScope := map[string]*model.Value{}
				providedExpression := "/dummy/deprecated/pkg-fs-ref"
				providedPkgRef := new(pkg.FakeHandle)

				fakeInterpolater := new(interpolater.Fake)
				// err to trigger immediate return
				fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

				objectUnderTest := _fileEvaluator{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.EvalToFile(
					providedScope,
					providedExpression,
					providedPkgRef,
					"dummyScratchDir",
				)

				/* assert */
				actualExpression,
					actualScope,
					actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

				Expect(actualExpression).To(Equal(providedExpression))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
			})
			Context("interpolater.Interpolate errors", func() {
				It("should return expected result", func() {
					/* arrange */
					providedExpression := "/deprecatedPkgFsRef"

					fakeInterpolater := new(interpolater.Fake)
					interpolateError := errors.New("dummyError")
					fakeInterpolater.InterpolateReturns("", interpolateError)

					expectedErr := fmt.Errorf(
						"unable to evaluate %v to file; error was %v",
						providedExpression,
						interpolateError.Error(),
					)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						providedExpression,
						new(pkg.FakeHandle),
						"dummyScratchDir",
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(expectedErr))
				})
			})
			Context("interpolater.Interpolate doesn't error", func() {
				It("should return expected result", func() {
					/* arrange */
					interpolatedExpression := "dummyExpression"

					fakeInterpolater := new(interpolater.Fake)
					fakeInterpolater.InterpolateReturns(interpolatedExpression, nil)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					fakeHandle := new(pkg.FakeHandle)
					pkgRef := "dummyPkgRef"
					fakeHandle.RefReturns(pkgRef)

					expectedFileValue := filepath.Join(pkgRef, interpolatedExpression)

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						"/deprecatedPkgFsRef",
						fakeHandle,
						"dummyScratchDir",
					)

					/* assert */
					Expect(*actualValue).To(Equal(model.Value{File: &expectedFileValue}))
					Expect(actualErr).To(BeNil())
				})
			})
		})
		Context("expression is dir scope ref w/ path", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				scopeName := "dummyScopeName"
				providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

				providedPath := "dummyPath"
				providedExpression := fmt.Sprintf("$(%v/%v)", scopeName, providedPath)
				providedPkgRef := new(pkg.FakeHandle)

				fakeInterpolater := new(interpolater.Fake)
				// err to trigger immediate return
				fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

				objectUnderTest := _fileEvaluator{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.EvalToFile(
					providedScope,
					providedExpression,
					providedPkgRef,
					"dummyScratchDir",
				)

				/* assert */
				actualExpression,
					actualScope,
					actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

				Expect(actualExpression).To(Equal(providedPath))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
			})
			Context("interpolater.Interpolate errs", func() {
				It("should return expected result", func() {

					/* arrange */
					scopeName := "dummyScopeName"
					providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

					providedPath := "dummyPath"
					providedExpression := fmt.Sprintf("$(%v/%v)", scopeName, providedPath)

					fakeInterpolater := new(interpolater.Fake)
					interpolateErr := errors.New("dummyError")
					fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

					expectedErr := fmt.Errorf(
						"unable to evaluate path %v; error was %v",
						providedPath,
						interpolateErr.Error(),
					)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						providedScope,
						providedExpression,
						new(pkg.FakeHandle),
						"dummyScratchDir",
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(expectedErr))

				})
			})
			Context("interpolater.Interpolate doesn't error", func() {
				It("should return expected result", func() {
					/* arrange */
					scopeName := "dummyScopeName"
					dirValue := "dummyDirValue"

					interpolatedPath := "dummyInterpolatedPath"
					fakeInterpolater := new(interpolater.Fake)
					fakeInterpolater.InterpolateReturns(interpolatedPath, nil)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					expectedValue := filepath.Join(dirValue, interpolatedPath)

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{scopeName: {Dir: &dirValue}},
						fmt.Sprintf("$(%v/path)", scopeName),
						new(pkg.FakeHandle),
						"dummyScratchDir",
					)

					/* assert */
					Expect(*actualValue).To(Equal(model.Value{File: &expectedValue}))
					Expect(actualErr).To(BeNil())
				})
			})
		})
		Context("expression is dir scope ref & deprecated path", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				scopeName := "dummyScopeName"
				providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

				providedPath := "/dummyPath"
				providedExpression := fmt.Sprintf("$(%v)%v", scopeName, providedPath)
				providedPkgRef := new(pkg.FakeHandle)

				fakeInterpolater := new(interpolater.Fake)
				// err to trigger immediate return
				fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

				objectUnderTest := _fileEvaluator{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.EvalToFile(
					providedScope,
					providedExpression,
					providedPkgRef,
					"dummyScratchDir",
				)

				/* assert */
				actualExpression,
					actualScope,
					actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

				Expect(actualExpression).To(Equal(providedPath))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
			})
			Context("interpolater.Interpolate errs", func() {
				It("should return expected result", func() {

					/* arrange */
					scopeName := "dummyScopeName"
					providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

					providedPath := "/dummyPath"
					providedExpression := fmt.Sprintf("$(%v)%v", scopeName, providedPath)

					fakeInterpolater := new(interpolater.Fake)
					interpolateErr := errors.New("dummyError")
					fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

					expectedErr := fmt.Errorf(
						"unable to evaluate path %v; error was %v",
						providedPath,
						interpolateErr.Error(),
					)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						providedScope,
						providedExpression,
						new(pkg.FakeHandle),
						"dummyScratchDir",
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(expectedErr))

				})
			})
			Context("interpolater.Interpolate doesn't error", func() {
				It("should return expected result", func() {
					/* arrange */
					scopeName := "dummyScopeName"
					dirValue := "dummyDirValue"
					providedScratchDir := "dummyScratchDir"

					fakeInterpolater := new(interpolater.Fake)
					interpolatedPath := "dummyInterpolatedPath"
					fakeInterpolater.InterpolateReturns(interpolatedPath, nil)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					expectedFileValue := filepath.Join(dirValue, interpolatedPath)

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{scopeName: {Dir: &dirValue}},
						fmt.Sprintf("$(%v)/path", scopeName),
						new(pkg.FakeHandle),
						providedScratchDir,
					)

					/* assert */
					Expect(*actualValue).To(Equal(model.Value{File: &expectedFileValue}))
					Expect(actualErr).To(BeNil())
				})
			})
		})
		Context("expression starts w/ scope ref", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				providedScope := map[string]*model.Value{"dummyKey": {String: new(string)}}
				providedExpression := "$(dummyRef)suffix"
				providedPkgHandle := new(pkg.FakeHandle)

				interpolatedPath := "dummyInterpolatedPath"
				fakeInterpolater := new(interpolater.Fake)
				// err to trigger immediate return
				fakeInterpolater.InterpolateReturns(interpolatedPath, errors.New("dummyError"))

				objectUnderTest := _fileEvaluator{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.EvalToFile(
					providedScope,
					providedExpression,
					providedPkgHandle,
					"dummyScratchDir",
				)

				/* assert */
				actualExpression,
					actualScope,
					actualPkgHandle := fakeInterpolater.InterpolateArgsForCall(0)

				Expect(actualExpression).To(Equal(providedExpression))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualPkgHandle).To(Equal(providedPkgHandle))
			})
			Context("interpolater.Interpolate errs", func() {
				It("should return expected result", func() {

					/* arrange */
					providedExpression := "$(dummyRef)suffix"

					fakeInterpolater := new(interpolater.Fake)
					interpolateErr := errors.New("dummyError")
					fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

					expectedErr := fmt.Errorf(
						"unable to evaluate %v to file; error was %v",
						providedExpression,
						interpolateErr.Error(),
					)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						providedExpression,
						new(pkg.FakeHandle),
						"dummyScratchDir",
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(expectedErr))

				})
			})
			Context("interpolater.Interpolate doesn't err", func() {
				It("should call data.CoerceToFile w/ expected args & return result", func() {
					/* arrange */
					providedExpression := "$(dummyRef)suffix"
					providedScratchDir := "dummyScratchDir"

					interpolatedExpression := "dummyInterpolatedExpression"
					fakeInterpolater := new(interpolater.Fake)
					fakeInterpolater.InterpolateReturns(interpolatedExpression, nil)

					fakeData := new(data.Fake)
					coerceValue := model.Value{File: new(string)}
					coerceErr := errors.New("dummyErr")
					fakeData.CoerceToFileReturns(&coerceValue, coerceErr)

					objectUnderTest := _fileEvaluator{
						data:         fakeData,
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						providedExpression,
						new(pkg.FakeHandle),
						providedScratchDir,
					)

					/* assert */
					actualCoerceValueArg, actualCoerceScratchDirArg := fakeData.CoerceToFileArgsForCall(0)
					Expect(*actualCoerceValueArg).To(Equal(model.Value{String: &interpolatedExpression}))
					Expect(actualCoerceScratchDirArg).To(Equal(providedScratchDir))

					Expect(*actualValue).To(Equal(coerceValue))
					Expect(actualErr).To(Equal(coerceErr))
				})
			})
		})
		Context("expression isn't ref", func() {
			It("should call interpolater.Interpolate w/ expected args", func() {
				/* arrange */
				providedScope := map[string]*model.Value{"dummyKey": {String: new(string)}}
				providedExpression := "dummyExpression"
				providedPkgHandle := new(pkg.FakeHandle)

				interpolatedPath := "dummyInterpolatedPath"
				fakeInterpolater := new(interpolater.Fake)
				// err to trigger immediate return
				fakeInterpolater.InterpolateReturns(interpolatedPath, errors.New("dummyError"))

				objectUnderTest := _fileEvaluator{
					interpolater: fakeInterpolater,
				}

				/* act */
				objectUnderTest.EvalToFile(
					providedScope,
					providedExpression,
					providedPkgHandle,
					"dummyScratchDir",
				)

				/* assert */
				actualExpression,
					actualScope,
					actualPkgHandle := fakeInterpolater.InterpolateArgsForCall(0)

				Expect(actualExpression).To(Equal(providedExpression))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualPkgHandle).To(Equal(providedPkgHandle))
			})
			Context("interpolater.Interpolate errs", func() {
				It("should return expected result", func() {

					/* arrange */
					providedExpression := "dummyExpression"

					fakeInterpolater := new(interpolater.Fake)
					interpolateErr := errors.New("dummyError")
					fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

					expectedErr := fmt.Errorf(
						"unable to evaluate %v to file; error was %v",
						providedExpression,
						interpolateErr.Error(),
					)

					objectUnderTest := _fileEvaluator{
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						providedExpression,
						new(pkg.FakeHandle),
						"dummyScratchDir",
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(expectedErr))

				})
			})
			Context("interpolater.Interpolate doesn't err", func() {
				It("should call data.CoerceToFile w/ expected args & return result", func() {
					/* arrange */
					providedExpression := "dummyExpression"
					providedScratchDir := "dummyScratchDir"

					interpolatedExpression := "dummyInterpolatedExpression"
					fakeInterpolater := new(interpolater.Fake)
					fakeInterpolater.InterpolateReturns(interpolatedExpression, nil)

					fakeData := new(data.Fake)
					coerceValue := model.Value{File: new(string)}
					coerceErr := errors.New("dummyErr")
					fakeData.CoerceToFileReturns(&coerceValue, coerceErr)

					objectUnderTest := _fileEvaluator{
						data:         fakeData,
						interpolater: fakeInterpolater,
					}

					/* act */
					actualValue, actualErr := objectUnderTest.EvalToFile(
						map[string]*model.Value{},
						providedExpression,
						new(pkg.FakeHandle),
						providedScratchDir,
					)

					/* assert */
					actualCoerceValueArg, actualCoerceScratchDirArg := fakeData.CoerceToFileArgsForCall(0)
					Expect(*actualCoerceValueArg).To(Equal(model.Value{String: &interpolatedExpression}))
					Expect(actualCoerceScratchDirArg).To(Equal(providedScratchDir))

					Expect(*actualValue).To(Equal(coerceValue))
					Expect(actualErr).To(Equal(coerceErr))
				})
			})
		})
	})
	Context("expression isnt float64, map[string]interface{}, or string", func() {
		It("should return expected result", func() {
			/* arrange */
			providedExpression := struct{}{}
			objectUnderTest := _fileEvaluator{}

			/* act */
			actualValue, actualErr := objectUnderTest.EvalToFile(
				map[string]*model.Value{},
				providedExpression,
				new(pkg.FakeHandle),
				"dummyScratchDir",
			)

			/* assert */
			Expect(actualValue).To(BeNil())
			Expect(actualErr).To(Equal(fmt.Errorf("unable to evaluate %+v to file", providedExpression)))
		})
	})
})