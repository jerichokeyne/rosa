package interactive_e2e

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/openshift/rosa/tests/ci/labels"
	"github.com/openshift/rosa/tests/utils/exec/rosacli"
	"github.com/openshift/rosa/tests/utils/helper"
	"github.com/openshift/rosa/tests/utils/interactive"
)

var _ = Describe("Account role tests", func() {
	It("Can create account roles with manual mode",
		labels.High, labels.Runtime.OCMResources, func() {
			rolePrefix := "asdf-test"
			By("Setup")
			session, err := interactive.NewInteractiveSession("rosa", "create", "account-roles", "-i")
			Expect(err).ToNot(HaveOccurred())
			defer session.Close()

			defaultTimeout := time.Second * 30

			By("Set role prefix")
			err = session.ExpectAndSend("Role prefix", rolePrefix, defaultTimeout)
			Expect(err).ToNot(HaveOccurred())
			By("Use default Permissions boundary ARN")
			err = session.ExpectAndSend("Permissions boundary ARN", "", defaultTimeout)
			Expect(err).ToNot(HaveOccurred())
			By("Use default Path")
			err = session.ExpectAndSend("Path", "", defaultTimeout)
			Expect(err).ToNot(HaveOccurred())
			By("Select manual creation mode with arrow keys")
			_, err = session.ExpectPrompt("Role creation mode", defaultTimeout)
			Expect(err).ToNot(HaveOccurred())
			err = session.Send(interactive.DOWN)
			Expect(err).ToNot(HaveOccurred())
			err = session.SendEnter()
			Expect(err).ToNot(HaveOccurred())
			By("Create classic roles")
			err = session.ExpectAndSend("Create Classic account roles", "y", defaultTimeout)
			Expect(err).ToNot(HaveOccurred())
			By("Create HCP roles")
			err = session.ExpectAndSend("Create Hosted CP account roles", "y", defaultTimeout)
			Expect(err).ToNot(HaveOccurred())
			err = session.ExpectAndSend("Use account roles for Hosted CP shared VPC", "n", defaultTimeout)
			Expect(err).ToNot(HaveOccurred())
			_, err = session.ExpectPrompt("All policy files saved to the current directory", defaultTimeout)
			Expect(err).ToNot(HaveOccurred())

			By("Run the AWS commands")
			outputBuf := bytes.NewBufferString(session.GetFullCleanOutput())
			commands := helper.ExtractCommandsToCreateAWSResources(*outputBuf)

			testAWSAccountID := "090777400063"
			fmt.Printf("Commands: %s\n", commands)
			for _, command := range commands {
				fmt.Printf("Running command: %s\n", command)
				command = strings.ReplaceAll(
					command, "arn:aws:iam::765374464689:",
					fmt.Sprintf("arn:aws:iam::%s:", testAWSAccountID))
				args := strings.Split(command, " ")
				cmd := exec.Command(args[0], args[1:]...)
				err = cmd.Run()
				Expect(err).To(BeNil())
			}

			session.DumpOutput()
			err = session.Wait(time.Second)
			Expect(err).ToNot(HaveOccurred())
			session.Close()

			rosaClient := rosacli.NewClient()
			ocmResourceService := rosaClient.OCMResource
			By("List the account roles created in manual mode")
			accountRoleList, _, err := ocmResourceService.ListAccountRole()
			Expect(err).To(BeNil())
			accountRoles := accountRoleList.AccountRoles(rolePrefix)
			Expect(len(accountRoles)).To(Equal(7))
			for _, ar := range accountRoles {
				Expect(ar.AWSManaged).To(Equal("Yes"))
			}

			By("Delete the account-roles in manual mode")
			_, err = ocmResourceService.DeleteAccountRole("--mode", "auto",
				"--prefix", rolePrefix,
				"-y")

			Expect(err).To(BeNil())
		})
})
